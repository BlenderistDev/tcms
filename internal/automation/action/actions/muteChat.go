package actions

import (
	"github.com/BlenderistDev/automation/datamapper"
	"github.com/BlenderistDev/automation/interfaces"
	interfaces2 "tcms/m/internal/automation/action/interfaces"
	"tcms/m/internal/model"
	"tcms/m/internal/telegramClient"
)

type muteChatAction struct {
	telegram telegramClient.TelegramClient
}

func CreateMuteChatAction(client telegramClient.TelegramClient) interfaces2.ActionWithModel {
	return muteChatAction{
		telegram: client,
	}
}

func (a muteChatAction) Execute(action model.Action, trigger interfaces.Trigger) error {
	dm := datamapper.DataMapper{Mapping: model.ConvertMappingToDmMapping(action.Mapping)}
	id, err := dm.GetFromMap(trigger.GetData(), "id")
	if err != nil {
		return err
	}

	unMute, err := dm.GetFromMapBool(trigger.GetData(), "unMute")
	if err != nil {
		return err
	}

	err = a.telegram.MuteChat(id, unMute)
	if err != nil {
		return err
	}

	return nil
}
