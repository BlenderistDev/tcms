package actions

import (
	interfaces2 "tcms/m/internal/action/interfaces"
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db/model"
	"tcms/m/internal/telegramClient"
)

type muteUserAction struct {
	telegram telegramClient.TelegramClient
}

func CreateMuteUserAction(client telegramClient.TelegramClient) interfaces2.ActionWithModel {
	return muteUserAction{
		telegram: client,
	}
}

func (a muteUserAction) Execute(action model.Action, trigger interfaces.Trigger) error {
	dm := datamapper.DataMapper{Mapping: model.ConvertMappingToDmMapping(action.Mapping)}
	peer, err := dm.GetFromMap(trigger.GetData(), "peer")
	if err != nil {
		return err
	}

	accessHash, err := dm.GetFromMap(trigger.GetData(), "accessHash")
	if err != nil {
		return err
	}

	unMute, err := dm.GetFromMapBool(trigger.GetData(), "unMute")
	if err != nil {
		return err
	}

	err = a.telegram.MuteUser(peer, accessHash, unMute)
	if err != nil {
		return err
	}

	return nil
}
