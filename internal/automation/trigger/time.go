package trigger

import (
	"time"

	"github.com/BlenderistDev/automation/interfaces"
	trigger2 "github.com/BlenderistDev/automation/trigger"
)

func StartTimeTrigger(triggerChan chan interfaces.TriggerEvent, d time.Duration) {
	ticker := time.NewTicker(d)
	go func() {
		for {
			select {
			case <-ticker.C:
				t := &trigger2.Trigger{}
				t.SetName("time")
				t.SetData("timestamp", time.Now().String())
				triggerChan <- t
			}
		}
	}()
}
