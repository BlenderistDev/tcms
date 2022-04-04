package action

import (
	"fmt"

	"tcms/m/internal/automation/action/actions"
	"tcms/m/internal/automation/action/interfaces"
	"tcms/m/internal/telegramClient"
	"tcms/m/pkg/tcms"
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
			ActionFields: []*tcms.Field{
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
			ActionFields: []*tcms.Field{
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
			ActionFields: []*tcms.Field{
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
