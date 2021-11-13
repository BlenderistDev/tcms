package core

import "fmt"

type Trigger interface {
	GetName() string
	GetData() interface{}
}

type Action interface {
	Execute(trigger Trigger) error
}

type Automation struct {
	Actions []Action
}

func (a Automation) Execute(trigger Trigger) {
	for _, action := range a.Actions {
		err := action.Execute(trigger)
		if err != nil {
			fmt.Println("Error while executing action")
			fmt.Println(err)
		}
	}
}
