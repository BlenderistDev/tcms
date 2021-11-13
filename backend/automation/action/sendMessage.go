package action

import (
	"fmt"
	"strconv"
	"tcms/m/automation/core"
	"tcms/m/db/model"
	"tcms/m/dry"
	"tcms/m/telegramClient"
)

type sendMessageAction struct {
	telegram *telegramClient.TelegramClient
	action   model.Action
}

func CreateSendMessageAction(action model.Action) core.Action {
	telegram := telegramClient.NewTelegram()
	return sendMessageAction{
		telegram: telegram,
		action:   action,
	}
}

func (a sendMessageAction) Execute(_ core.Trigger) {
	peer, err := a.getFromMapInt("peer")
	if err != nil {
		dry.HandleError(err)
	}

	accessHash, err := a.getFromMapInt("accessHash")
	if err != nil {
		dry.HandleError(err)
	}

	message, err := a.getFromMap("message")
	if err != nil {
		dry.HandleError(err)
	}

	err = a.telegram.SendMessage(message, int32(peer), int64(accessHash))
	if err != nil {
		fmt.Printf("Send message error")
	}
}

func (a sendMessageAction) getFromMapInt(key string) (int, error) {
	s, err := a.getFromMap(key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (a sendMessageAction) getFromMap(key string) (string, error) {
	mappingData, ok := a.action.Mapping[key]
	if ok {
		return mappingData.Value, nil
	}
	return "", fmt.Errorf("key %s not found", key)
}
