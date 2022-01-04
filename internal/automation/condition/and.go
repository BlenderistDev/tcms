package condition

import (
	"fmt"
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/automation/interfaces"
)

type andCondition struct {
	datamapper.DataMapper
	subConditions []interfaces.Condition
}

func createAndCondition(dataMapper datamapper.DataMapper, subConditions []interfaces.Condition) (interfaces.Condition, error) {
	if len(subConditions) < 2 {
		return nil, fmt.Errorf("and condition should have at least two subconditions")
	}

	return andCondition{
		DataMapper:    dataMapper,
		subConditions: subConditions,
	}, nil
}

func (c andCondition) Check(trigger interfaces.Trigger) (bool, error) {
	res := true
	for _, subCondition := range c.subConditions {
		subRes, err := subCondition.Check(trigger)
		if err != nil {
			return false, err
		}
		res = res && subRes
	}
	return res, nil
}
