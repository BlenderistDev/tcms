package action

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"tcms/internal/automation/action/interfaces"
	telegramClient2 "tcms/internal/testing/telegramClient"
)

func TestFactory_CreateAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tg := telegramClient2.NewMockTelegramClient(ctrl)
	actionFactory := NewFactory(tg)

	action, err := actionFactory.CreateAction("sendMessage")
	assert.Nil(t, err)
	switch action.(type) {
	case interfaces.ActionWithModel:
	default:
		t.Errorf("action sendMessage is not created")
	}
}

func TestCreateAction_createMuteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tg := telegramClient2.NewMockTelegramClient(ctrl)
	actionFactory := NewFactory(tg)

	action, err := actionFactory.CreateAction("muteUser")
	assert.Nil(t, err)
	switch action.(type) {
	case interfaces.ActionWithModel:
	default:
		t.Errorf("action sendMessage is not created")
	}
}

func TestCreateAction_createMuteChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tg := telegramClient2.NewMockTelegramClient(ctrl)
	actionFactory := NewFactory(tg)

	action, err := actionFactory.CreateAction("muteChat")
	assert.Nil(t, err)
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

	tg := telegramClient2.NewMockTelegramClient(ctrl)
	actionFactory := NewFactory(tg)

	_, err := actionFactory.CreateAction(name)
	assert.Equal(t, "unknown action "+name, err.Error())
}

func TestGetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tg := telegramClient2.NewMockTelegramClient(ctrl)
	actionFactory := NewFactory(tg)
	actionFactory.GetList()
}
