package action

import (
	"reflect"
	"tcms/m/db/model"
	"testing"
)

func TestCreateSendMessageAction(t *testing.T) {
	actionModel := model.Action{
		Name:    "name",
		Mapping: nil,
	}
	createdAction := CreateSendMessageAction(actionModel)

	switch action := createdAction.(type) {
	case sendMessageAction:
		if !reflect.DeepEqual(action.DataMapper.Action, actionModel) {
			t.Errorf("expected %v, got %v", actionModel, action.DataMapper.Action)
		}
	default:
		t.Errorf("action type is not sendMessageAction")
	}
}
