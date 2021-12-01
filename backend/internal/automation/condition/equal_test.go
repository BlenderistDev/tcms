package condition

import (
	"tcms/m/internal/db/model"
	"tcms/m/internal/dry"
	"testing"
)

func TestCreateSendMessageAction(t *testing.T) {
	conditionModel := &model.Condition{
		Name:    "name",
		Mapping: nil,
	}
	createdCondition := createEqualCondition(conditionModel)

	switch condition := createdCondition.(type) {
	case equalCondition:
		dry.TestCheckEqual(t, conditionModel, condition.DataMapper.Condition)
	default:
		t.Errorf("condition type is not sendMessageAction")
	}
}
