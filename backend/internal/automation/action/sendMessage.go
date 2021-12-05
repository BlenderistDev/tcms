package action

import (
	"fmt"
	"tcms/m/internal/automation/core"
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/db/model"
	"tcms/m/internal/telegramClient"
)

type sendMessageAction struct {
	telegram telegramClient.TelegramClient
	datamapper.DataMapper
}

func createSendMessageAction(action model.Action) core.Action {
	return sendMessageAction{
		telegram: telegramClient.NewTelegram(),
		DataMapper: datamapper.DataMapper{
			Mapping: action.Mapping,
		},
	}
}

func (a sendMessageAction) Execute(trigger core.Trigger) error {
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
