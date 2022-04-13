package actions

import (
	"fmt"
	"testing"

	mock_interfaces "github.com/BlenderistDev/automation/testing/interfaces"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"tcms/internal/dry"
	"tcms/internal/model"
	telegramClient2 "tcms/internal/testing/telegramClient"
)

func TestCreateSendMessageAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	createdAction := CreateSendMessageAction(telegramClient)

	switch createdAction.(type) {
	case sendMessageAction:
	default:
		t.Errorf("action type is not sendMessageAction")
	}
}

func TestSendMessageAction_Execute(t *testing.T) {
	const (
		messageKey   = "message"
		messageValue = "test message"
		peerKey      = "peer"
		peerValue    = "123123"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)
	telegramClient.
		EXPECT().
		SendMessage(gomock.Eq(peerValue), gomock.Eq(messageValue))

	actionModel := model.Action{
		Name: "name",
		Mapping: map[string]model.Mapping{
			peerKey: {
				Simple: true,
				Name:   peerKey,
				Value:  peerValue,
			},
			messageKey: {
				Simple: true,
				Name:   messageKey,
				Value:  messageValue,
			},
		},
	}

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	trigger.EXPECT().GetData().Return(make(map[string]string))
	trigger.EXPECT().GetData().Return(make(map[string]string))

	sendMessageAction := CreateSendMessageAction(telegramClient)
	err := sendMessageAction.Execute(actionModel, trigger)
	assert.Nil(t, err)
}

func TestSendMessageAction_Execute_peerError(t *testing.T) {
	const (
		messageKey      = "message"
		messageValue    = "test message"
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
			messageKey: {
				Simple: true,
				Name:   messageKey,
				Value:  messageValue,
			},
		},
	}

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	trigger.EXPECT().GetData().Return(make(map[string]string))

	sendMessageAction := CreateSendMessageAction(telegramClient)
	err := sendMessageAction.Execute(actionModel, trigger)
	dry.TestCheckEqual(t, "key peer not found", err.Error())
}

func TestSendMessageAction_Execute_messageError(t *testing.T) {
	const (
		peerKey   = "peer"
		peerValue = "123123"
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

	sendMessageAction := CreateSendMessageAction(telegramClient)
	err := sendMessageAction.Execute(actionModel, trigger)
	dry.TestCheckEqual(t, "key message not found", err.Error())
}

func TestSendMessageAction_Execute_telegramError(t *testing.T) {
	const (
		messageKey   = "message"
		messageValue = "test message"
		peerKey      = "peer"
		peerValue    = "123123"
		errorTexxt   = "some error"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)
	telegramClient.
		EXPECT().
		SendMessage(gomock.Eq(peerValue), gomock.Eq(messageValue)).
		Return(fmt.Errorf(errorTexxt))

	actionModel := model.Action{
		Name: "name",
		Mapping: map[string]model.Mapping{
			peerKey: {
				Simple: true,
				Name:   peerKey,
				Value:  peerValue,
			},
			messageKey: {
				Simple: true,
				Name:   messageKey,
				Value:  messageValue,
			},
		},
	}

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	trigger.EXPECT().GetData().Return(make(map[string]string))
	trigger.EXPECT().GetData().Return(make(map[string]string))

	sendMessageAction := CreateSendMessageAction(telegramClient)
	err := sendMessageAction.Execute(actionModel, trigger)
	dry.TestCheckEqual(t, errorTexxt, err.Error())
}
