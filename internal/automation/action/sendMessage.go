package action

import (
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db/model"
	"tcms/m/internal/telegramClient"
)

type sendMessageAction struct {
	telegram telegramClient.TelegramClient
	datamapper.DataMapper
}

func createSendMessageAction(action model.Action, client telegramClient.TelegramClient) interfaces.Action {
	return sendMessageAction{
		telegram: client,
		DataMapper: datamapper.DataMapper{
			Mapping: action.Mapping,
		},
	}
}

func (a sendMessageAction) Execute(trigger interfaces.Trigger) error {
	peer, err := a.GetFromMap(trigger, "peer")
	if err != nil {
		return err
	}

	message, err := a.GetFromMap(trigger, "message")
	if err != nil {
		return err
	}

	err = a.telegram.SendMessage(peer, message)
	if err != nil {
		return err
	}

	return nil
}
