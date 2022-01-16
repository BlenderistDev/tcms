package automation

import (
	"fmt"
	"tcms/m/internal/automation/core"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/dry"
)

type Service struct {
	list map[string][]core.Automation
}

// Start launch automation service
func (s *Service) Start(triggerChan chan interfaces.Trigger) {
	for {
		trigger := <-triggerChan
		s.HandleTrigger(trigger)
	}
}

func (s *Service) Init() {
	s.list = make(map[string][]core.Automation)
}

func (s *Service) AddAutomation(automation core.Automation) {
	for _, trigger := range automation.GetTriggers() {
		s.list[trigger] = append(s.list[trigger], automation)
	}
}

func (s *Service) HandleTrigger(trigger interfaces.Trigger) {
	automationList := s.list[trigger.GetName()]
	if automationList == nil {
		fmt.Printf("no automation for trigger %s\n", trigger.GetName())
		return
	}
	for _, automation := range automationList {
		fmt.Printf("Trigger with type %s\n", trigger.GetName())
		err := automation.Execute(trigger)
		dry.HandleError(err)
	}
}
