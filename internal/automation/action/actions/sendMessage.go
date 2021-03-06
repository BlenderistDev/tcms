package actions

import (
	"github.com/BlenderistDev/automation/datamapper"
	"github.com/BlenderistDev/automation/interfaces"
	interfaces2 "tcms/internal/automation/action/interfaces"
	"tcms/internal/model"
	"tcms/internal/telegramClient"
)

type sendMessageAction struct {
	telegram telegramClient.TelegramClient
}

func CreateSendMessageAction(client telegramClient.TelegramClient) interfaces2.ActionWithModel {
	return sendMessageAction{
		telegram: client,
	}
}

func (a sendMessageAction) Execute(action model.Action, trigger interfaces.TriggerEvent) error {
	dm := datamapper.DataMapper{Mapping: model.ConvertMappingToDmMapping(action.Mapping)}
	peer, err := dm.GetFromMap(trigger.GetData(), "peer")
	if err != nil {
		return err
	}

	message, err := dm.GetFromMap(trigger.GetData(), "message")
	if err != nil {
		return err
	}

	err = a.telegram.SendMessage(peer, message)
	if err != nil {
		return err
	}

	return nil
}
