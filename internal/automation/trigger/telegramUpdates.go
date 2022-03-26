package trigger

import (
	"encoding/json"

	"github.com/BlenderistDev/automation/interfaces"
	"github.com/sirupsen/logrus"
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

func StartTelegramUpdateTrigger(addConsumer chan chan []uint8, triggerChan chan interfaces.TriggerEvent, log *logrus.Logger) {
	ch := make(chan []uint8)
	addConsumer <- ch
	for {
		data := <-ch
		var trigger telegramUpdateTrigger
		err := json.Unmarshal(data, &trigger)
		if err == nil {
			triggerChan <- trigger
		} else {
			log.Info(err)
		}

	}
}
