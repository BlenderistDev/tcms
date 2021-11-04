package automation

import "fmt"

type Trigger interface {
	GetName() string
	GetData() interface{}
}

func HandleTrigger(trigger Trigger) {
	fmt.Printf("Trigger with type %s\n", trigger.GetName())
}
