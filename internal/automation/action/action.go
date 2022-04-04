package action

import (
	"github.com/BlenderistDev/automation/interfaces"
	interfaces2 "tcms/internal/automation/action/interfaces"
	"tcms/internal/model"
)

type actionWithModel struct {
	action interfaces2.ActionWithModel
	model  model.Action
}

func GetActionWithModel(action interfaces2.ActionWithModel, model model.Action) interfaces.Action {
	return &actionWithModel{
		action: action,
		model:  model,
	}
}

func (m *actionWithModel) Execute(trigger interfaces.TriggerEvent) error {
	return m.action.Execute(m.model, trigger)
}
