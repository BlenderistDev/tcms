package action

import (
	"fmt"
	"tcms/m/automation/core"
	"tcms/m/db/model"
	"tcms/m/telegramClient"
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
	peer, err := a.getFromMapInt(trigger, "peer")
	if err != nil {
		return err
	}

	accessHash, err := a.getFromMapInt(trigger, "accessHash")
	if err != nil {
		return err
	}

	message, err := a.getFromMap(trigger, "message")
	if err != nil {
		return err
	}

	err = a.telegram.SendMessage(message, int32(peer), int64(accessHash))
	if err != nil {
		return fmt.Errorf("send message error")
	}

	return nil
}
