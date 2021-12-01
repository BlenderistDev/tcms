package action

import (
	"fmt"
	"tcms/m/internal/automation/core"
	"tcms/m/internal/db/model"
	"tcms/m/internal/telegramClient"
)

type sendMessageAction struct {
	telegram telegramClient.TelegramClient
	DataMapper
}

func createSendMessageAction(action model.Action) core.Action {
	return sendMessageAction{
		telegram: telegramClient.NewTelegram(),
		DataMapper: DataMapper{
			Action: action,
		},
	}
}

func (a sendMessageAction) Execute(trigger core.Trigger) error {
	peer, err := a.getFromMapInt32(trigger, "peer")
	if err != nil {
		return err
	}

	accessHash, err := a.getFromMapInt64(trigger, "accessHash")
	if err != nil {
		return err
	}

	message, err := a.getFromMap(trigger, "message")
	if err != nil {
		return err
	}

	err = a.telegram.SendMessage(message, peer, accessHash)
	if err != nil {
		return fmt.Errorf("send message error")
	}

	return nil
}
