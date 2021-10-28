package automation

import "fmt"

type Trigger interface {
	GetName() string
	GetKeyList() []string
	GetData() map[string]string
}

func HandleTrigger(trigger Trigger) {
	fmt.Printf("Trigger with type %s", trigger.GetName())
}
