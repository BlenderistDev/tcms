package condition

import (
	"tcms/m/db/model"
	"tcms/m/dry"
	"testing"
)

func TestCreateCondition_createEqual(t *testing.T) {
	conditionModel := model.Condition{
		Name:    "equal",
		Mapping: nil,
	}
	condition, err := CreateCondition(&conditionModel)
	dry.HandleError(err)
	switch condition.(type) {
	case equalCondition:
	default:
		t.Errorf("condition type is not equal")
	}
}
