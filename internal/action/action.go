package action

import (
	"github.com/BlenderistDev/automation/interfaces"
	interfaces2 "tcms/m/internal/action/interfaces"
	"tcms/m/internal/db/model"
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

func (m *actionWithModel) Execute(trigger interfaces.Trigger) error {
	return m.action.Execute(m.model, trigger)
}
