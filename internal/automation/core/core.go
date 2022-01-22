package core

import (
	"fmt"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/dry"
)

type Automation struct {
	actions   []interfaces.Action
	condition interfaces.Condition
	triggers  []string
}

func (a *Automation) Execute(trigger interfaces.Trigger) error {
	if a.condition == nil || a.checkCondition(trigger) {
		err := a.executeActions(trigger)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Automation) AddTrigger(trigger string) {
	a.triggers = append(a.triggers, trigger)
}

func (a *Automation) GetTriggers() []string {
	return a.triggers
}

func (a *Automation) AddAction(action interfaces.Action) {
	a.actions = append(a.actions, action)
}

func (a *Automation) AddCondition(condition interfaces.Condition) {
	a.condition = condition
}

func (a *Automation) checkCondition(trigger interfaces.Trigger) bool {
	res, err := a.condition.Check(trigger)
	if err != nil {
		dry.HandleError(err)
		return false
	}
	return res
}

func (a *Automation) executeActions(trigger interfaces.Trigger) error {
	for _, action := range a.actions {
		err := action.Execute(trigger)
		if err != nil {
			return fmt.Errorf("error while executing action: %s", err.Error())
		}
	}
	return nil
}
