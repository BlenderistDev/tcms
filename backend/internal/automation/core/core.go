package core

import (
	"fmt"
	"tcms/m/internal/dry"
)

type Trigger interface {
	GetName() string
	GetData() map[string]string
}

type Action interface {
	Execute(trigger Trigger) error
}

type Condition interface {
	Check(trigger Trigger) (bool, error)
}

type Automation struct {
	Actions   []Action
	Condition Condition
}

func (a Automation) Execute(trigger Trigger) {
	if a.Condition == nil || a.checkCondition(trigger) {
		a.executeActions(trigger)
	}
}

func (a Automation) checkCondition(trigger Trigger) bool {
	res, err := a.Condition.Check(trigger)
	if err != nil {
		dry.HandleError(err)
		return false
	}
	return res
}

func (a Automation) executeActions(trigger Trigger) {
	for _, action := range a.Actions {
		err := action.Execute(trigger)
		if err != nil {
			fmt.Println("Error while executing action")
			fmt.Println(err)
		}
	}
}
