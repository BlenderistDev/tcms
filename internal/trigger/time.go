package trigger

import (
	"tcms/m/internal/automation/interfaces"
	"time"
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

func StartTimeTrigger(triggerChan chan interfaces.Trigger) {
	ticker := time.NewTicker(time.Second)
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
