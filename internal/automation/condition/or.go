package condition

import (
	"fmt"
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/automation/interfaces"
)

type orCondition struct {
	datamapper.DataMapper
	subConditions []interfaces.Condition
}

func createOrCondition(dataMapper datamapper.DataMapper, subConditions []interfaces.Condition) (interfaces.Condition, error) {
	if len(subConditions) < 2 {
		return nil, fmt.Errorf("or condition should have at least two subconditions")
	}

	return orCondition{
		DataMapper:    dataMapper,
		subConditions: subConditions,
	}, nil
}

func (c orCondition) Check(trigger interfaces.Trigger) (bool, error) {
	res := false
	for _, subCondition := range c.subConditions {
		subRes, err := subCondition.Check(trigger)
		if err != nil {
			return false, err
		}
		res = res || subRes
	}
	return res, nil
}
