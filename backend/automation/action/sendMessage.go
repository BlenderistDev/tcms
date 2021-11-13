package action

import (
	"fmt"
	"strconv"
	"tcms/m/automation/core"
	"tcms/m/db/model"
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

func (a sendMessageAction) getFromMapInt(trigger core.Trigger, key string) (int, error) {
	s, err := a.getFromMap(trigger, key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (a sendMessageAction) getFromMap(trigger core.Trigger, key string) (string, error) {
	mappingData, ok := a.action.Mapping[key]
	if ok {
		if mappingData.Simple {
			return mappingData.Value, nil
		} else {
			triggerData := trigger.GetData()
			value, ok := triggerData[mappingData.Value]
			if ok {
				return value, nil
			} else {
				return "", fmt.Errorf("key %s not found in trigger data", key)
			}
		}

	}
	return "", fmt.Errorf("key %s not found", key)
}
