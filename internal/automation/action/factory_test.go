package action

import (
	"github.com/golang/mock/gomock"
	"tcms/m/internal/db/model"
	"tcms/m/internal/dry"
	telegramClient2 "tcms/m/internal/testing/telegramClient"
	"testing"
)

func TestCreateAction_createSendMessage(t *testing.T) {
	actionModel := model.Action{
		Name:    "sendMessage",
		Mapping: nil,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	action, err := CreateAction(actionModel, telegramClient)
	dry.TestHandleError(t, err)
	switch action.(type) {
	case sendMessageAction:
	default:
		t.Errorf("action type is not sendMessageAction")
	}
}

func TestCreateAction_unknownAction(t *testing.T) {
	const name = "someAction"
	actionModel := model.Action{
		Name:    name,
		Mapping: nil,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	_, err := CreateAction(actionModel, telegramClient)
	dry.TestCheckEqual(t, "unknown action "+name, err.Error())
}
