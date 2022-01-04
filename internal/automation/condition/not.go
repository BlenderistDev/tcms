package condition

import (
	"fmt"
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db/model"
)

type notCondition struct {
	datamapper.DataMapper
	subCondition interfaces.Condition
}

func createNotCondition(condition *model.Condition) (interfaces.Condition, error) {
	if len(condition.SubConditions) != 1 {
		return nil, fmt.Errorf("not condition can have only one subcondition")
	}
	subCondition, err := CreateCondition(&condition.SubConditions[0])
	if err != nil {
		return nil, err
	}
	return notCondition{
		DataMapper:   datamapper.DataMapper{Mapping: condition.Mapping},
		subCondition: subCondition,
	}, nil
}

func (c notCondition) Check(trigger interfaces.Trigger) (bool, error) {
	res, err := c.subCondition.Check(trigger)
	if err != nil {
		return false, err
	}
	return !res, nil
}
