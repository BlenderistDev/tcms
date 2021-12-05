package action

import (
	"tcms/m/internal/db/model"
	"tcms/m/internal/dry"
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
	createdAction := createSendMessageAction(actionModel)

	switch action := createdAction.(type) {
	case sendMessageAction:
		dry.TestCheckEqual(t, actionModel.Mapping, action.DataMapper.Mapping)
	default:
		t.Errorf("action type is not sendMessageAction")
	}
}

func TestSendMessageAction_Execute(t *testing.T) {

}
