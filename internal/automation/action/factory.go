package action

import (
	"fmt"
	"tcms/m/internal/telegramClient"
)

func CreateAction(name string, telegram telegramClient.TelegramClient) (Action, error) {
	var action Action
	switch name {
	case "sendMessage":
		action = CreateSendMessageAction(telegram)
	case "muteUser":
		action = CreateMuteUserAction(telegram)
	case "muteChat":
		action = CreateMuteChatAction(telegram)
	default:
		return nil, fmt.Errorf("unknown action %s", name)
	}
	return action, nil
}
