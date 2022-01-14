package core

import (
	"fmt"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db/model"
	"tcms/m/internal/dry"
)

type Action interface {
	Execute(action model.Action, trigger interfaces.Trigger) error
}

type ActionWithModel struct {
	Action interfaces.Action
	Model  model.Action
}

func (m ActionWithModel) Execute(trigger interfaces.Trigger) error {
	return m.Action.Execute(m.Model, trigger)
}

type Automation struct {
	Actions   []ActionWithModel
	Condition interfaces.Condition
}

func (a Automation) Execute(trigger interfaces.Trigger) error {
	if a.Condition == nil || a.checkCondition(trigger) {
		err := a.executeActions(trigger)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a Automation) checkCondition(trigger interfaces.Trigger) bool {
	res, err := a.Condition.Check(trigger)
	if err != nil {
		dry.HandleError(err)
		return false
	}
	return res
}

func (a Automation) executeActions(trigger interfaces.Trigger) error {
	for _, action := range a.Actions {
		err := action.Execute(trigger)
		if err != nil {
			return fmt.Errorf("error while executing action: %s", err.Error())
		}
	}
	return nil
}
