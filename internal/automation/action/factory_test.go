package action

import (
	"testing"

	"github.com/golang/mock/gomock"
	"tcms/internal/automation/action/interfaces"
	"tcms/internal/dry"
	telegramClient2 "tcms/internal/testing/telegramClient"
)

func TestCreateAction_createSendMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	action, err := CreateAction("sendMessage", telegramClient)
	dry.TestHandleError(t, err)
	switch action.(type) {
	case interfaces.ActionWithModel:
	default:
		t.Errorf("action sendMessage is not created")
	}
}

func TestCreateAction_createMuteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	action, err := CreateAction("muteUser", telegramClient)
	dry.TestHandleError(t, err)
	switch action.(type) {
	case interfaces.ActionWithModel:
	default:
		t.Errorf("action sendMessage is not created")
	}
}

func TestCreateAction_createMuteChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	action, err := CreateAction("muteChat", telegramClient)
	dry.TestHandleError(t, err)
	switch action.(type) {
	case interfaces.ActionWithModel:
	default:
		t.Errorf("action sendMessage is not created")
	}
}

func TestCreateAction_unknownAction(t *testing.T) {
	const name = "someAction"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	_, err := CreateAction(name, telegramClient)
	dry.TestCheckEqual(t, "unknown action "+name, err.Error())
}

func TestGetList(t *testing.T) {
	GetList()
}
