package action

import (
	"github.com/golang/mock/gomock"
	"tcms/m/internal/dry"
	telegramClient2 "tcms/m/internal/testing/telegramClient"
	"testing"
)

func TestCreateAction_createSendMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	action, err := CreateAction("sendMessage", telegramClient)
	dry.TestHandleError(t, err)
	switch action.(type) {
	case sendMessageAction:
	default:
		t.Errorf("action type is not sendMessageAction")
	}
}

func TestCreateAction_createMuteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	action, err := CreateAction("muteUser", telegramClient)
	dry.TestHandleError(t, err)
	switch action.(type) {
	case muteUserAction:
	default:
		t.Errorf("action type is not muteUserAction")
	}
}

func TestCreateAction_createMuteChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	telegramClient := telegramClient2.NewMockTelegramClient(ctrl)

	action, err := CreateAction("muteChat", telegramClient)
	dry.TestHandleError(t, err)
	switch action.(type) {
	case muteChatAction:
	default:
		t.Errorf("action type is not muteUserAction")
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
