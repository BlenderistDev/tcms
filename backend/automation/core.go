package automation

import "fmt"

type Trigger interface {
	GetName() string
	GetData() interface{}
}

type Action interface {
	Execute(trigger Trigger)
}

func HandleTrigger(trigger Trigger) {
	fmt.Printf("Trigger with type %s\n", trigger.GetName())
}
