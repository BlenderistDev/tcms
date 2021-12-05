package action

import (
	"github.com/golang/mock/gomock"
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

}
