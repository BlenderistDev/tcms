package condition

import (
	"testing"

	"github.com/BlenderistDev/automation/interfaces"
	"github.com/stretchr/testify/assert"
	"tcms/internal/dry"
	"tcms/internal/model"
)

func TestCreateCondition_createEqual(t *testing.T) {
	conditionModel := model.Condition{
		Name:    "equal",
		Mapping: nil,
	}
	conditionFactory := NewFactory()
	condition, err := conditionFactory.CreateCondition(&conditionModel)
	assert.Nil(t, err)
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
	conditionFactory := NewFactory()
	_, err := conditionFactory.CreateCondition(&conditionModel)
	dry.TestCheckEqual(t, "unknown condition "+name, err.Error())
}

func TestGetList(t *testing.T) {
	conditionFactory := NewFactory()
	conditionFactory.GetList()
}
