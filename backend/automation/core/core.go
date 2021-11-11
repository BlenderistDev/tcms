package core

type Trigger interface {
	GetName() string
	GetData() interface{}
}

type Action interface {
	Execute(trigger Trigger)
}

type Automation struct {
	Action Action
}

func (a Automation) Execute(trigger Trigger) {
	a.Action.Execute(trigger)
}
