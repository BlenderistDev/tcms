package action

import (
	"fmt"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db/model"
	"tcms/m/internal/telegramClient"
)

func CreateAction(actionData model.Action) (interfaces.Action, error) {
	telegram, err := telegramClient.NewTelegram()
	if err != nil {
		return nil, err
	}
	var action interfaces.Action
	switch actionData.Name {
	case "sendMessage":
		action = createSendMessageAction(actionData, telegram)
	default:
		return nil, fmt.Errorf("unknown action %s", actionData.Name)
	}
	return action, nil
}
