package action

import (
	"fmt"

	"tcms/internal/automation/action/actions"
	"tcms/internal/automation/action/interfaces"
	"tcms/internal/telegramClient"
	"tcms/pkg/tcms"
)

type Factory interface {
	CreateAction(name string) (interfaces.ActionWithModel, error)
	GetList() *tcms.ActionList
}

type factory struct {
	tg telegramClient.TelegramClient
}

func NewFactory(tg telegramClient.TelegramClient) Factory {
	return factory{tg: tg}
}

func (f factory) CreateAction(name string) (interfaces.ActionWithModel, error) {
	var action interfaces.ActionWithModel
	switch name {
	case "sendMessage":
		action = actions.CreateSendMessageAction(f.tg)
	case "muteUser":
		action = actions.CreateMuteUserAction(f.tg)
	case "muteChat":
		action = actions.CreateMuteChatAction(f.tg)
	default:
		return nil, fmt.Errorf("unknown action %s", name)
	}
	return action, nil
}

func (f factory) GetList() *tcms.ActionList {
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
