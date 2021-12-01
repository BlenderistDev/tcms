package action

import (
	"tcms/m/internal/db/model"
	"tcms/m/internal/dry"
	"testing"
)

func TestCreateSendMessageAction(t *testing.T) {
	actionModel := model.Action{
		Name:    "name",
		Mapping: nil,
	}
	createdAction := createSendMessageAction(actionModel)

	switch action := createdAction.(type) {
	case sendMessageAction:
		dry.TestCheckEqual(t, actionModel, action.DataMapper.Action)
	default:
		t.Errorf("action type is not sendMessageAction")
	}
}
