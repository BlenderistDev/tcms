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

type actionWithModel struct {
	action interfaces.Action
	model  model.Action
}

func GetActionWithModel(action interfaces.Action, model model.Action) interfaces.ActionWithModel {
	return &actionWithModel{
		action: action,
		model:  model,
	}
}

func (m *actionWithModel) Execute(trigger interfaces.Trigger) error {
	return m.action.Execute(m.model, trigger)
}

type Automation struct {
	Actions   []interfaces.ActionWithModel
	Condition interfaces.Condition
	Triggers  []string
}

func (a *Automation) Execute(trigger interfaces.Trigger) error {
	if a.Condition == nil || a.checkCondition(trigger) {
		err := a.executeActions(trigger)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Automation) AddAction(action interfaces.ActionWithModel) {
	a.Actions = append(a.Actions, action)
}

func (a Automation) AddCondition(condition interfaces.Condition) {
	a.Condition = condition
}

func (a *Automation) checkCondition(trigger interfaces.Trigger) bool {
	res, err := a.Condition.Check(trigger)
	if err != nil {
		dry.HandleError(err)
		return false
	}
	return res
}

func (a *Automation) executeActions(trigger interfaces.Trigger) error {
	for _, action := range a.Actions {
		err := action.Execute(trigger)
		if err != nil {
			return fmt.Errorf("error while executing action: %s", err.Error())
		}
	}
	return nil
}
