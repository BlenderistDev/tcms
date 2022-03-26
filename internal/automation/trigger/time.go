package trigger

import (
	"time"

	"github.com/BlenderistDev/automation/interfaces"
)

type timeTrigger struct {
	Name string
	Data map[string]string
}

func (t timeTrigger) GetName() string {
	return t.Name
}

func (t timeTrigger) GetData() map[string]string {
	return t.Data
}

func StartTimeTrigger(triggerChan chan interfaces.TriggerEvent, d time.Duration) {
	ticker := time.NewTicker(d)
	go func() {
		for {
			select {
			case <-ticker.C:
				trigger := timeTrigger{
					Name: "time",
					Data: map[string]string{
						"timestamp": time.Now().String(),
					},
				}
				triggerChan <- trigger
			}
		}
	}()
}
