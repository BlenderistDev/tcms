package action

import (
	"tcms/m/internal/db/model"
	"tcms/m/internal/dry"
	"testing"
)

func TestCreateAction_createSendMessage(t *testing.T) {
	actionModel := model.Action{
		Name:    "sendMessage",
		Mapping: nil,
	}
	action, err := CreateAction(actionModel)
	dry.HandleError(err)
	switch action.(type) {
	case sendMessageAction:
	default:
		t.Errorf("action type is not sendMessageAction")
	}
}
