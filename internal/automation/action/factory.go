package action

import (
	"fmt"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db/model"
	"tcms/m/internal/telegramClient"
)

func CreateAction(actionData model.Action, telegram telegramClient.TelegramClient) (interfaces.Action, error) {
	var action interfaces.Action
	switch actionData.Name {
	case "sendMessage":
		action = CreateSendMessageAction(telegram)
	case "muteUser":
		action = CreateMuteUserAction(telegram)
	case "muteChat":
		action = CreateMuteChatAction(telegram)
	default:
		return nil, fmt.Errorf("unknown action %s", actionData.Name)
	}
	return action, nil
}
