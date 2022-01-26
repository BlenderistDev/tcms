package condition

import (
	"github.com/BlenderistDev/automation/datamapper"
	"github.com/BlenderistDev/automation/interfaces"
)

type equalCondition struct {
	dm datamapper.DataMapper
}

func CreateEqualCondition(dataMapper datamapper.DataMapper) interfaces.Condition {
	return equalCondition{
		dm: dataMapper,
	}
}

func (c equalCondition) Check(trigger interfaces.Trigger) (bool, error) {
	value1, err := c.dm.GetFromMap(trigger.GetData(), "value1")
	if err != nil {
		return false, err
	}
	value2, err := c.dm.GetFromMap(trigger.GetData(), "value2")
	if err != nil {
		return false, err
	}
	return value1 == value2, nil
}
