package condition

import (
	"testing"

	"github.com/BlenderistDev/automation/interfaces"
	"tcms/m/internal/dry"
	"tcms/m/internal/model"
)

func TestCreateCondition_createEqual(t *testing.T) {
	conditionModel := model.Condition{
		Name:    "equal",
		Mapping: nil,
	}
	condition, err := CreateCondition(&conditionModel)
	dry.TestHandleError(t, err)
	switch condition.(type) {
	case interfaces.Condition:
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
	dry.TestCheckEqual(t, "unknown condition "+name, err.Error())
}

func TestGetList(t *testing.T) {
	GetList()
}
