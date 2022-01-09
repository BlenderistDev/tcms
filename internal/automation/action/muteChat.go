package action

import (
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db/model"
	"tcms/m/internal/telegramClient"
)

type muteChatAction struct {
	telegram telegramClient.TelegramClient
	datamapper.DataMapper
}

func createMuteChatAction(action model.Action, client telegramClient.TelegramClient) interfaces.Action {
	return muteChatAction{
		telegram: client,
		DataMapper: datamapper.DataMapper{
			Mapping: action.Mapping,
		},
	}
}

func (a muteChatAction) Execute(trigger interfaces.Trigger) error {
	id, err := a.GetFromMap(trigger, "id")
	if err != nil {
		return err
	}

	err = a.telegram.MuteChat(id, true)
	if err != nil {
		return err
	}

	return nil
}
