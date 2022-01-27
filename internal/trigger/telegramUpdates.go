package trigger

import (
	"encoding/json"
	"github.com/BlenderistDev/automation/interfaces"
	"tcms/m/internal/dry"
)

type telegramUpdateTrigger struct {
	Name string            `json:"name"`
	Data map[string]string `json:"data"`
}

func (t telegramUpdateTrigger) GetName() string {
	return t.Name
}

func (t telegramUpdateTrigger) GetData() map[string]string {
	return t.Data
}

func StartTelegramTrigger(addConsumer chan chan []uint8, triggerChan chan interfaces.Trigger) {
	ch := make(chan []uint8)
	addConsumer <- ch
	for {
		data := <-ch
		var trigger telegramUpdateTrigger
		err := json.Unmarshal(data, &trigger)
		dry.HandleError(err)
		triggerChan <- trigger
	}
}
