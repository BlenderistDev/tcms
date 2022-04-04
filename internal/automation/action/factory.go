package action

import (
	"fmt"

	"tcms/internal/automation/action/actions"
	"tcms/internal/automation/action/interfaces"
	"tcms/internal/telegramClient"
	"tcms/pkg/tcms"
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

func GetList() *tcms.ActionList {
	return &tcms.ActionList{Actions: []*tcms.ActionDescription{
		{
			Name: "sendMessage",
			Fields: []*tcms.Field{
				{
					Name:     "peer",
					Type:     "string",
					Required: true,
				},
				{
					Name:     "message",
					Type:     "string",
					Required: true,
				},
			},
		},
		{
			Name: "muteUser",
			Fields: []*tcms.Field{
				{
					Name:     "peer",
					Type:     "string",
					Required: true,
				},
				{
					Name:     "accessHash",
					Type:     "string",
					Required: true,
				},
				{
					Name:     "onMute",
					Type:     "bool",
					Required: true,
				},
			},
		},
		{
			Name: "muteChat",
			Fields: []*tcms.Field{
				{
					Name:     "id",
					Type:     "string",
					Required: true,
				},
				{
					Name:     "unMute",
					Type:     "bool",
					Required: true,
				},
			},
		},
	}}
}
