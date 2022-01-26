package interfaces

import (
	"github.com/BlenderistDev/automation/interfaces"
	"tcms/m/internal/db/model"
)

type ActionWithModel interface {
	Execute(action model.Action, trigger interfaces.Trigger) error
}
