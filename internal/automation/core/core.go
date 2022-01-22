package core

import (
	"fmt"
	"tcms/m/internal/automation/interfaces"
)

type automation struct {
	actions   []interfaces.Action
	condition interfaces.Condition
	triggers  []string
}

func GetAutomation() interfaces.Automation {
	return &automation{}
}

func (a *automation) Execute(trigger interfaces.Trigger) error {
	if a.condition != nil {
		checkRes, err := a.checkCondition(trigger)
		if err != nil {
			return err
		}
		if !checkRes {
			return nil
		}
	}
	err := a.executeActions(trigger)
	if err != nil {
		return err
	}
	return nil
}

func (a *automation) AddTrigger(trigger string) {
	a.triggers = append(a.triggers, trigger)
}

func (a *automation) GetTriggers() []string {
	return a.triggers
}

func (a *automation) AddAction(action interfaces.Action) {
	a.actions = append(a.actions, action)
}

func (a *automation) AddCondition(condition interfaces.Condition) {
	a.condition = condition
}

func (a *automation) checkCondition(trigger interfaces.Trigger) (bool, error) {
	res, err := a.condition.Check(trigger)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (a *automation) executeActions(trigger interfaces.Trigger) error {
	for _, action := range a.actions {
		err := action.Execute(trigger)
		if err != nil {
			return fmt.Errorf("error while executing action: %s", err.Error())
		}
	}
	return nil
}
