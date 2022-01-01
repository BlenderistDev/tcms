package automation

import (
	"encoding/json"
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

func UpdateTriggerFactory(addConsumer chan chan []uint8) {
	ch := make(chan []uint8)
	addConsumer <- ch

	automationService := Service{}
	automationService.Start()

	for {
		data := <-ch
		var trigger TelegramUpdateTrigger
		err := json.Unmarshal(data, &trigger)
		dry.HandleError(err)
		automationService.HandleTrigger(trigger)
	}
}
