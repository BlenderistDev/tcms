package condition

import (
	"tcms/m/internal/automation/interfaces"
)

type notCondition struct {
	subCondition interfaces.Condition
}

func createNotCondition(condition interfaces.Condition) interfaces.Condition {
	return notCondition{
		subCondition: condition,
	}
}

func (c notCondition) Check(trigger interfaces.Trigger) (bool, error) {
	res, err := c.subCondition.Check(trigger)
	if err != nil {
		return false, err
	}
	return !res, nil
}
