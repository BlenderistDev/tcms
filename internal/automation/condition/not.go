package condition

import (
	"fmt"
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/automation/interfaces"
)

type notCondition struct {
	subCondition interfaces.Condition
}

func createNotCondition(_ datamapper.DataMapper, subConditions []interfaces.Condition) (interfaces.Condition, error) {
	if len(subConditions) != 1 {
		return nil, fmt.Errorf("not condition can have only one subcondition")
	}

	return notCondition{
		subCondition: subConditions[0],
	}, nil
}

func (c notCondition) Check(trigger interfaces.Trigger) (bool, error) {
	res, err := c.subCondition.Check(trigger)
	if err != nil {
		return false, err
	}
	return !res, nil
}
