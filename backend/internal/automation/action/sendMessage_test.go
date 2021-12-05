package action

import (
	"github.com/golang/mock/gomock"
	"tcms/m/internal/automation/core"
	"tcms/m/internal/db/model"
	"tcms/m/internal/dry"
	telegramClient2 "tcms/m/internal/telegramClient/mock"
	"testing"
)

func TestCreateSendMessageAction(t *testing.T) {
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

	createdAction := createSendMessageAction(actionModel, telegramClient)

	switch action := createdAction.(type) {
	case sendMessageAction:
		dry.TestCheckEqual(t, actionModel.Mapping, action.DataMapper.Mapping)
	default:
		t.Errorf("action type is not sendMessageAction")
	}
}

func TestSendMessageAction_Execute(t *testing.T) {
	const (
		messageKey      = "message"
		messageValue    = "test message"
		accessHashKey   = "accessHash"
		accessHashValue = "456456"
		peerKey         = "peer"
		peerValue       = "123123"
	)

	var peerValueInt int32 = 123123
	var accessHashValueInt int64 = 456456

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)
	telegramClient.EXPECT().SendMessage(gomock.Eq(messageValue), gomock.Eq(peerValueInt), gomock.Eq(accessHashValueInt))

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
			messageKey: {
				Simple: true,
				Name:   messageKey,
				Value:  messageValue,
			},
		},
	}

	trigger := core.NewMockTrigger(ctrl)
	sendMessageAction := createSendMessageAction(actionModel, telegramClient)
	err := sendMessageAction.Execute(trigger)
	dry.TestHandleError(t, err)
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

	trigger := core.NewMockTrigger(ctrl)
	sendMessageAction := createSendMessageAction(actionModel, telegramClient)
	err := sendMessageAction.Execute(trigger)
	dry.TestCheckEqual(t, "key peer not found", err.Error())
}

func TestSendMessageAction_Execute_accessHashError(t *testing.T) {
	const (
		messageKey   = "message"
		messageValue = "test message"
		peerKey      = "peer"
		peerValue    = "123123"
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
			messageKey: {
				Simple: true,
				Name:   messageKey,
				Value:  messageValue,
			},
		},
	}

	trigger := core.NewMockTrigger(ctrl)
	sendMessageAction := createSendMessageAction(actionModel, telegramClient)
	err := sendMessageAction.Execute(trigger)
	dry.TestCheckEqual(t, "key accessHash not found", err.Error())
}

func TestSendMessageAction_Execute_messageError(t *testing.T) {
	const (
		accessHashKey   = "accessHash"
		accessHashValue = "456456"
		peerKey         = "peer"
		peerValue       = "123123"
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

	trigger := core.NewMockTrigger(ctrl)
	sendMessageAction := createSendMessageAction(actionModel, telegramClient)
	err := sendMessageAction.Execute(trigger)
	dry.TestCheckEqual(t, "key message not found", err.Error())
}
