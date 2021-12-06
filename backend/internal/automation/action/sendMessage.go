package action

import (
	"fmt"
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
	peer, err := a.GetFromMapInt32(trigger, "peer")
	if err != nil {
		return err
	}

	accessHash, err := a.GetFromMapInt64(trigger, "accessHash")
	if err != nil {
		return err
	}

	message, err := a.GetFromMap(trigger, "message")
	if err != nil {
		return err
	}

	err = a.telegram.SendMessage(message, peer, accessHash)
	if err != nil {
		return fmt.Errorf("send message error")
	}

	return nil
}
