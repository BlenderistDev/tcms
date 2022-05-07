package trigger

import (
	"encoding/json"

	"github.com/BlenderistDev/automation/interfaces"
	trigger2 "github.com/BlenderistDev/automation/trigger"
	"github.com/sirupsen/logrus"
)

type telegramUpdate struct {
	Name string            `json:"name"`
	Data map[string]string `json:"data"`
}

func StartTelegramUpdateTrigger(addConsumer chan chan []uint8, triggerChan chan interfaces.TriggerEvent, log *logrus.Logger) {
	ch := make(chan []uint8)
	addConsumer <- ch
	for {
		data := <-ch
		var trigger telegramUpdate
		err := json.Unmarshal(data, &trigger)

		t := makeFromTelegramUpdate(trigger)

		if err == nil {
			triggerChan <- t
		} else {
			log.Info(err)
		}
	}
}

func makeFromTelegramUpdate(tg telegramUpdate) *trigger2.Trigger {
	t := &trigger2.Trigger{}
	t.SetName(tg.Name)

	for k, v := range tg.Data {
		t.SetData(k, v)
	}

	return t
}
