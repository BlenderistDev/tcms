package interfaces

type Trigger interface {
	GetName() string
	GetData() map[string]string
}

type Condition interface {
	Check(trigger Trigger) (bool, error)
}

type Action interface {
	Execute(trigger Trigger) error
}
