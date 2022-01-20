package condition

import (
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/automation/interfaces"
)

type equalCondition struct {
	datamapper.DataMapper
}

func createEqualCondition(dataMapper datamapper.DataMapper, _ []interfaces.Condition) interfaces.Condition {
	return equalCondition{
		DataMapper: dataMapper,
	}
}

func (c equalCondition) Check(trigger interfaces.Trigger) (bool, error) {
	value1, err := c.GetFromMap(trigger.GetData(), "value1")
	if err != nil {
		return false, err
	}
	value2, err := c.GetFromMap(trigger.GetData(), "value2")
	if err != nil {
		return false, err
	}
	return value1 == value2, nil
}
