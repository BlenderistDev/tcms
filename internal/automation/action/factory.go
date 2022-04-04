package action

import (
	"fmt"

	"tcms/m/internal/automation/action/actions"
	"tcms/m/internal/automation/action/interfaces"
	"tcms/m/internal/telegramClient"
)

func CreateAction(name string, telegram telegramClient.TelegramClient) (interfaces.ActionWithModel, error) {
	var action interfaces.ActionWithModel
	switch name {
	case "sendMessage":
		action = actions.CreateSendMessageAction(telegram)
	case "muteUser":
		action = actions.CreateMuteUserAction(telegram)
	case "muteChat":
		action = actions.CreateMuteChatAction(telegram)
	default:
		return nil, fmt.Errorf("unknown action %s", name)
	}
	return action, nil
}
