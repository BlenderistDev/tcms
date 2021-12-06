package core

import (
	"fmt"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/dry"
)

type Automation struct {
	Actions   []interfaces.Action
	Condition interfaces.Condition
}

func (a Automation) Execute(trigger interfaces.Trigger) {
	if a.Condition == nil || a.checkCondition(trigger) {
		a.executeActions(trigger)
	}
}

func (a Automation) checkCondition(trigger interfaces.Trigger) bool {
	res, err := a.Condition.Check(trigger)
	if err != nil {
		dry.HandleError(err)
		return false
	}
	return res
}

func (a Automation) executeActions(trigger interfaces.Trigger) {
	for _, action := range a.Actions {
		err := action.Execute(trigger)
		if err != nil {
			fmt.Println("Error while executing action")
			fmt.Println(err)
		}
	}
}
