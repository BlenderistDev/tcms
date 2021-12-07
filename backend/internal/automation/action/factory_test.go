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
	_, err := CreateAction(actionModel)
	dry.TestCheckEqual(t, "unknown action "+name, err.Error())
}
