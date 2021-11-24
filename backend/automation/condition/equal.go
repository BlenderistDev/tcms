package condition

import (
	"tcms/m/automation/core"
	"tcms/m/db/model"
)

type equalCondition struct {
	DataMapper
}

func CreateEqualCondition(condition model.Condition) core.Condition {
	return equalCondition{
		DataMapper: DataMapper{Condition: condition},
	}
}

func (c equalCondition) Check(trigger core.Trigger) (bool, error) {
	value1, err := c.getFromMap(trigger, "value1")
	if err != nil {
		return false, err
	}
	value2, err := c.getFromMap(trigger, "value2")
	if err != nil {
		return false, err
	}
	return value1 == value2, nil
}
