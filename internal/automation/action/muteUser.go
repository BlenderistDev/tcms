package action

import (
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db/model"
	"tcms/m/internal/telegramClient"
)

type muteUserAction struct {
	telegram telegramClient.TelegramClient
	datamapper.DataMapper
}

func createMuteUserAction(action model.Action, client telegramClient.TelegramClient) interfaces.Action {
	return muteUserAction{
		telegram: client,
		DataMapper: datamapper.DataMapper{
			Mapping: action.Mapping,
		},
	}
}

func (a muteUserAction) Execute(trigger interfaces.Trigger) error {
	peer, err := a.GetFromMap(trigger, "peer")
	if err != nil {
		return err
	}

	accessHash, err := a.GetFromMap(trigger, "accessHash")
	if err != nil {
		return err
	}

	err = a.telegram.MuteUser(peer, accessHash)
	if err != nil {
		return err
	}

	return nil
}
