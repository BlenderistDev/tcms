package action

import (
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db/model"
	"tcms/m/internal/telegramClient"
)

type muteUserAction struct {
	telegram telegramClient.TelegramClient
}

func CreateMuteUserAction(client telegramClient.TelegramClient) Action {
	return muteUserAction{
		telegram: client,
	}
}

func (a muteUserAction) Execute(action model.Action, trigger interfaces.Trigger) error {
	dm := datamapper.DataMapper{Mapping: model.ConvertMappingToDmMapping(action.Mapping)}
	peer, err := dm.GetFromMap(trigger, "peer")
	if err != nil {
		return err
	}

	accessHash, err := dm.GetFromMap(trigger, "accessHash")
	if err != nil {
		return err
	}

	unMute, err := dm.GetFromMapBool(trigger, "unMute")
	if err != nil {
		return err
	}

	err = a.telegram.MuteUser(peer, accessHash, unMute)
	if err != nil {
		return err
	}

	return nil
}
