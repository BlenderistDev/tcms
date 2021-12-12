package condition

import (
	"tcms/m/internal/db/model"
	"tcms/m/internal/dry"
	"testing"
)

func TestCreateCondition_createEqual(t *testing.T) {
	conditionModel := model.Condition{
		Name:    "equal",
		Mapping: nil,
	}
	condition, err := CreateCondition(&conditionModel)
	dry.TestHandleError(t, err)
	switch condition.(type) {
	case equalCondition:
	default:
		t.Errorf("condition type is not equal")
	}
}

func TestCreateCondition_NoCondition(t *testing.T) {
	const name = "notExistConditionName"
	conditionModel := model.Condition{
		Name: name,
	}
	_, err := CreateCondition(&conditionModel)
	dry.TestCheckEqual(t, "unknown action "+name, err.Error())
}
