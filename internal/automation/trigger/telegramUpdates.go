package trigger

import (
	"encoding/json"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/dry"
)

type TelegramUpdateTrigger struct {
	Name string            `json:"name"`
	Data map[string]string `json:"data"`
}

func (t TelegramUpdateTrigger) GetName() string {
	return t.Name
}

func (t TelegramUpdateTrigger) GetData() map[string]string {
	return t.Data
}

func StartTelegramTrigger(addConsumer chan chan []uint8, triggerChan chan interfaces.Trigger) {
	ch := make(chan []uint8)
	addConsumer <- ch
	for {
		data := <-ch
		var trigger TelegramUpdateTrigger
		err := json.Unmarshal(data, &trigger)
		dry.HandleError(err)
		triggerChan <- trigger
	}
}
