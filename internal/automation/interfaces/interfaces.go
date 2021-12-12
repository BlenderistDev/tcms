package interfaces

type Trigger interface {
	GetName() string
	GetData() map[string]string
}

type Action interface {
	Execute(trigger Trigger) error
}

type Condition interface {
	Check(trigger Trigger) (bool, error)
}
