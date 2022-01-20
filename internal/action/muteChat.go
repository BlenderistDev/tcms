package action

import (
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db/model"
	"tcms/m/internal/telegramClient"
)

type muteChatAction struct {
	telegram telegramClient.TelegramClient
}

func CreateMuteChatAction(client telegramClient.TelegramClient) Action {
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
