package action

import (
	"github.com/golang/mock/gomock"
	"tcms/m/internal/db/model"
	"tcms/m/internal/dry"
	mock_interfaces "tcms/m/internal/testing/automation/interfaces"
	telegramClient2 "tcms/m/internal/testing/telegramClient"
	"testing"
)

func TestCreateMuteChatAction(t *testing.T) {
	actionModel := model.Action{
		Name: "name",
		Mapping: map[string]model.Mapping{
			"test": {
				Simple: true,
				Name:   "name",
				Value:  "value",
			}},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	createdAction := createMuteChatAction(actionModel, telegramClient)

	switch action := createdAction.(type) {
	case muteChatAction:
		dry.TestCheckEqual(t, actionModel.Mapping, action.DataMapper.Mapping)
	default:
		t.Errorf("action type is not muteChatAction")
	}
}

func TestMuteChatAction_Execute_idError(t *testing.T) {
	const (
		unMuteKey   = "unMute"
		unMuteValue = ""
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	actionModel := model.Action{
		Name: "name",
		Mapping: map[string]model.Mapping{
			unMuteKey: {
				Simple: true,
				Name:   unMuteKey,
				Value:  unMuteValue,
			},
		},
	}

	trigger := mock_interfaces.NewMockTrigger(ctrl)
	muteUserAction := createMuteChatAction(actionModel, telegramClient)
	err := muteUserAction.Execute(trigger)
	dry.TestCheckEqual(t, "key id not found", err.Error())
}

func TestMuteChatAction_Execute_unMuteError(t *testing.T) {
	const (
		idKey   = "id"
		idValue = "456456"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	actionModel := model.Action{
		Name: "name",
		Mapping: map[string]model.Mapping{
			idKey: {
				Simple: true,
				Name:   idKey,
				Value:  idValue,
			},
		},
	}

	trigger := mock_interfaces.NewMockTrigger(ctrl)
	muteUserAction := createMuteChatAction(actionModel, telegramClient)
	err := muteUserAction.Execute(trigger)
	dry.TestCheckEqual(t, "key unMute not found", err.Error())
}
