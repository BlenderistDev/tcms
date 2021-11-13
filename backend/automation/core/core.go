package core

type Trigger interface {
	GetName() string
	GetData() interface{}
}

type Action interface {
	Execute(trigger Trigger)
}

type Automation struct {
	Actions []Action
}

func (a Automation) Execute(trigger Trigger) {
	for _, action := range a.Actions {
		action.Execute(trigger)
	}
}
