package action

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"tcms/m/internal/db/model"
	"tcms/m/internal/dry"
	mock_interfaces "tcms/m/internal/testing/automation/interfaces"
	telegramClient2 "tcms/m/internal/testing/telegramClient"
	"testing"
)

func TestCreateMuteUserAction(t *testing.T) {
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

	createdAction := createMuteUserAction(actionModel, telegramClient)

	switch action := createdAction.(type) {
	case muteUserAction:
		dry.TestCheckEqual(t, actionModel.Mapping, action.DataMapper.Mapping)
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

	trigger := mock_interfaces.NewMockTrigger(ctrl)
	muteUserAction := createMuteUserAction(actionModel, telegramClient)
	err := muteUserAction.Execute(trigger)
	dry.TestCheckEqual(t, "key peer not found", err.Error())
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

	trigger := mock_interfaces.NewMockTrigger(ctrl)
	muteUserAction := createMuteUserAction(actionModel, telegramClient)
	err := muteUserAction.Execute(trigger)
	dry.TestCheckEqual(t, "key accessHash not found", err.Error())
}

func TestMuteUserAction_Execute(t *testing.T) {
	const (
		accessHashKey   = "accessHash"
		accessHashValue = "456456"
		peerKey         = "peer"
		peerValue       = "456456"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)
	telegramClient.
		EXPECT().
		MuteUser(gomock.Eq(peerValue), gomock.Eq(accessHashValue))

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

	trigger := mock_interfaces.NewMockTrigger(ctrl)
	muteUserAction := createMuteUserAction(actionModel, telegramClient)
	err := muteUserAction.Execute(trigger)
	dry.TestHandleError(t, err)
}

func TestMuteUserAction_Execute_telegramError(t *testing.T) {
	const (
		accessHashKey   = "accessHash"
		accessHashValue = "456456"
		peerKey         = "peer"
		peerValue       = "456456"
		errorText       = "some error"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)
	telegramClient.
		EXPECT().
		MuteUser(gomock.Eq(peerValue), gomock.Eq(accessHashValue)).
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
		},
	}

	trigger := mock_interfaces.NewMockTrigger(ctrl)
	muteUserAction := createMuteUserAction(actionModel, telegramClient)
	err := muteUserAction.Execute(trigger)
	dry.TestCheckEqual(t, errorText, err.Error())
}
