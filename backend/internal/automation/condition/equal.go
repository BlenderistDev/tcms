package condition

import (
	"tcms/m/internal/automation/core"
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/db/model"
)

type equalCondition struct {
	datamapper.DataMapper
}

func createEqualCondition(condition *model.Condition) core.Condition {
	return equalCondition{
		DataMapper: datamapper.DataMapper{Mapping: condition.Mapping},
	}
}

func (c equalCondition) Check(trigger core.Trigger) (bool, error) {
	value1, err := c.GetFromMap(trigger, "value1")
	if err != nil {
		return false, err
	}
	value2, err := c.GetFromMap(trigger, "value2")
	if err != nil {
		return false, err
	}
	return value1 == value2, nil
}
