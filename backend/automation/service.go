package automation

import (
	"fmt"
	action2 "tcms/m/automation/action"
	"tcms/m/automation/core"
)

type Service struct {
	list map[string][]core.Automation
}

// Start launch automation service
func (s *Service) Start() {
	s.list = make(map[string][]core.Automation, 0)
	UpdateUserStatusList := make([]core.Automation, 0)
	action := action2.CreateSendMessageAction()
	s.list["UpdateUserStatus"] = append(UpdateUserStatusList, core.Automation{Action: action})
}

func (s *Service) HandleTrigger(trigger core.Trigger) {
	automationList := s.list[trigger.GetName()]
	if automationList == nil {
		fmt.Println("no automation")
		return
	}
	for _, automation := range automationList {
		fmt.Printf("Trigger with type %s\n", trigger.GetName())
		automation.Execute(trigger)
	}
}
