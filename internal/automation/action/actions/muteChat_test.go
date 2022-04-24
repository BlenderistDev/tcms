package actions

import (
	"fmt"
	"testing"

	mock_interfaces "github.com/BlenderistDev/automation/testing/interfaces"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"tcms/internal/model"
	telegramClient2 "tcms/internal/testing/telegramClient"
)

func TestCreateMuteChatAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	createdAction := CreateMuteChatAction(telegramClient)

	switch createdAction.(type) {
	case muteChatAction:
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

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	trigger.EXPECT().GetData().Return(make(map[string]string))

	muteUserAction := CreateMuteChatAction(telegramClient)
	err := muteUserAction.Execute(actionModel, trigger)
	assert.Equal(t, "key id not found", err.Error())
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

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	trigger.EXPECT().GetData().Return(make(map[string]string))
	trigger.EXPECT().GetData().Return(make(map[string]string))

	muteUserAction := CreateMuteChatAction(telegramClient)
	err := muteUserAction.Execute(actionModel, trigger)
	assert.Equal(t, "key unMute not found", err.Error())
}

func TestMuteChatAction_Execute(t *testing.T) {
	const (
		idKey       = "id"
		idValue     = "456456"
		unMuteKey   = "unMute"
		unMuteValue = ""
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)
	telegramClient.
		EXPECT().
		MuteChat(gomock.Eq(idValue), gomock.Eq(false))

	actionModel := model.Action{
		Name: "name",
		Mapping: map[string]model.Mapping{
			idKey: {
				Simple: true,
				Name:   idKey,
				Value:  idValue,
			},
			unMuteKey: {
				Simple: true,
				Name:   unMuteKey,
				Value:  unMuteValue,
			},
		},
	}

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	trigger.EXPECT().GetData().Return(make(map[string]string))
	trigger.EXPECT().GetData().Return(make(map[string]string))

	muteUserAction := CreateMuteChatAction(telegramClient)
	err := muteUserAction.Execute(actionModel, trigger)
	assert.Nil(t, err)
}

func TestMuteChatAction_Execute_telegramError(t *testing.T) {
	const (
		idKey       = "id"
		idValue     = "456456"
		unMuteKey   = "unMute"
		unMuteValue = ""
		errorText   = "some error"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)
	telegramClient.
		EXPECT().
		MuteChat(gomock.Eq(idValue), gomock.Eq(false)).
		Return(fmt.Errorf(errorText))

	actionModel := model.Action{
		Name: "name",
		Mapping: map[string]model.Mapping{
			idKey: {
				Simple: true,
				Name:   idKey,
				Value:  idValue,
			},
			unMuteKey: {
				Simple: true,
				Name:   unMuteKey,
				Value:  unMuteValue,
			},
		},
	}

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	trigger.EXPECT().GetData().Return(make(map[string]string))
	trigger.EXPECT().GetData().Return(make(map[string]string))

	muteUserAction := CreateMuteChatAction(telegramClient)
	err := muteUserAction.Execute(actionModel, trigger)
	assert.Equal(t, errorText, err.Error())
}
