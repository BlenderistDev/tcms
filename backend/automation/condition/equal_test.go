package condition

import (
	"reflect"
	"tcms/m/db/model"
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
		if !reflect.DeepEqual(condition.DataMapper.Condition, conditionModel) {
			t.Errorf("expected %v, got %v", conditionModel, condition.DataMapper.Condition)
		}
	default:
		t.Errorf("condition type is not sendMessageAction")
	}
}
