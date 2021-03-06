package interfaces

import (
	"github.com/BlenderistDev/automation/interfaces"
	"tcms/internal/model"
)

type ActionWithModel interface {
	Execute(action model.Action, trigger interfaces.TriggerEvent) error
}
