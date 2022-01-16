package action

import (
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db/model"
)

type Action interface {
	Execute(action model.Action, trigger interfaces.Trigger) error
}

type actionWithModel struct {
	action Action
	model  model.Action
}

func GetActionWithModel(action Action, model model.Action) interfaces.ActionWithModel {
	return &actionWithModel{
		action: action,
		model:  model,
	}
}

func (m *actionWithModel) Execute(trigger interfaces.Trigger) error {
	return m.action.Execute(m.model, trigger)
}
