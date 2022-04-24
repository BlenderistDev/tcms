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

func TestCreateMuteUserAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	createdAction := CreateMuteUserAction(telegramClient)

	switch createdAction.(type) {
	case muteUserAction:
	default:
		t.Errorf("action type is not muteUserAction")
	}
}

func TestMuteUserAction_Execute_peerError(t *testing.T) {
	const (
		accessHashKey   = "accessHash"
		accessHashValue = "456456"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	actionModel := model.Action{
		Name: "name",
		Mapping: map[string]model.Mapping{
			accessHashKey: {
				Simple: true,
				Name:   accessHashKey,
				Value:  accessHashValue,
			},
		},
	}

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	trigger.EXPECT().GetData().Return(make(map[string]string))

	muteUserAction := CreateMuteUserAction(telegramClient)
	err := muteUserAction.Execute(actionModel, trigger)
	assert.Equal(t, "key peer not found", err.Error())
}

func TestMuteUserAction_Execute_accessHashError(t *testing.T) {
	const (
		peerKey   = "peer"
		peerValue = "456456"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	actionModel := model.Action{
		Name: "name",
		Mapping: map[string]model.Mapping{
			peerKey: {
				Simple: true,
				Name:   peerKey,
				Value:  peerValue,
			},
		},
	}

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	trigger.EXPECT().GetData().Return(make(map[string]string))
	trigger.EXPECT().GetData().Return(make(map[string]string))

	muteUserAction := CreateMuteUserAction(telegramClient)
	err := muteUserAction.Execute(actionModel, trigger)
	assert.Equal(t, "key accessHash not found", err.Error())
}

func TestMuteUserAction_Execute_unMuteError(t *testing.T) {
	const (
		accessHashKey   = "accessHash"
		accessHashValue = "456456"
		peerKey         = "peer"
		peerValue       = "456456"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	actionModel := model.Action{
		Name: "name",
		Mapping: map[string]model.Mapping{
			peerKey: {
				Simple: true,
				Name:   peerKey,
				Value:  peerValue,
			},
			accessHashKey: {
				Simple: true,
				Name:   accessHashKey,
				Value:  accessHashValue,
			},
		},
	}

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	trigger.EXPECT().GetData().Return(make(map[string]string))
	trigger.EXPECT().GetData().Return(make(map[string]string))
	trigger.EXPECT().GetData().Return(make(map[string]string))

	muteUserAction := CreateMuteUserAction(telegramClient)
	err := muteUserAction.Execute(actionModel, trigger)
	assert.Equal(t, "key unMute not found", err.Error())
}

func TestMuteUserAction_Execute(t *testing.T) {
	const (
		accessHashKey   = "accessHash"
		accessHashValue = "456456"
		peerKey         = "peer"
		peerValue       = "456456"
		unMuteKey       = "unMute"
		unMuteValue     = ""
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)
	telegramClient.
		EXPECT().
		MuteUser(gomock.Eq(peerValue), gomock.Eq(accessHashValue), gomock.Eq(false))

	actionModel := model.Action{
		Name: "name",
		Mapping: map[string]model.Mapping{
			peerKey: {
				Simple: true,
				Name:   peerKey,
				Value:  peerValue,
			},
			accessHashKey: {
				Simple: true,
				Name:   accessHashKey,
				Value:  accessHashValue,
			},
			unMuteKey: {
				Simple: true,
				Name:   unMuteKey,
				Value:  unMuteValue,
			},
		},
	}

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	trigger.EXPECT().GetData().Return(make(map[string]string)).Times(3)

	muteUserAction := CreateMuteUserAction(telegramClient)
	err := muteUserAction.Execute(actionModel, trigger)
	assert.Nil(t, err)
}

func TestMuteUserAction_Execute_telegramError(t *testing.T) {
	const (
		accessHashKey   = "accessHash"
		accessHashValue = "456456"
		peerKey         = "peer"
		peerValue       = "456456"
		errorText       = "some error"
		unMuteKey       = "unMute"
		unMuteValue     = ""
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)
	telegramClient.
		EXPECT().
		MuteUser(gomock.Eq(peerValue), gomock.Eq(accessHashValue), gomock.Eq(false)).
		Return(fmt.Errorf(errorText))

	actionModel := model.Action{
		Name: "name",
		Mapping: map[string]model.Mapping{
			peerKey: {
				Simple: true,
				Name:   peerKey,
				Value:  peerValue,
			},
			accessHashKey: {
				Simple: true,
				Name:   accessHashKey,
				Value:  accessHashValue,
			},
			unMuteKey: {
				Simple: true,
				Name:   unMuteKey,
				Value:  unMuteValue,
			},
		},
	}

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	trigger.EXPECT().GetData().Return(make(map[string]string)).Times(3)

	muteUserAction := CreateMuteUserAction(telegramClient)
	err := muteUserAction.Execute(actionModel, trigger)
	assert.Equal(t, errorText, err.Error())
}
