package trigger

import (
	"tcms/m/internal/automation/interfaces"
	"time"
)

type TimeTrigger struct {
	Name string
	Data map[string]string
}

func (t TimeTrigger) GetName() string {
	return t.Name
}

func (t TimeTrigger) GetData() map[string]string {
	return t.Data
}

func StartTimeTrigger(triggerChan chan interfaces.Trigger) {
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				trigger := TimeTrigger{
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
