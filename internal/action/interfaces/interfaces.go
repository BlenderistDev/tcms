package interfaces

import (
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db/model"
)

type ActionWithModel interface {
	Execute(action model.Action, trigger interfaces.Trigger) error
}
