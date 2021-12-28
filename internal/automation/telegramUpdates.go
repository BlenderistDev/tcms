package automation

import (
	"encoding/json"
	"tcms/m/internal/dry"
	"tcms/m/internal/kafka"
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

func UpdateTriggerFactory() {
	client, err := kafka.NewKafkaClient()
	dry.HandleError(err)

	ch := make(chan []uint8)
	go client.Subscribe(ch)

	automationService := Service{}
	automationService.Start()

	for {
		data := <-ch
		var trigger TelegramUpdateTrigger
		err = json.Unmarshal(data, &trigger)
		dry.HandleError(err)
		automationService.HandleTrigger(trigger)
	}
}
