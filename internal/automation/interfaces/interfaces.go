package interfaces

import "tcms/m/internal/db/model"

type Trigger interface {
	GetName() string
	GetData() map[string]string
}

type Action interface {
	Execute(action model.Action, trigger Trigger) error
}

type Condition interface {
	Check(trigger Trigger) (bool, error)
}
