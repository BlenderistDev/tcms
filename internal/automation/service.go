package automation

import (
	"fmt"
	"tcms/m/internal/automation/core"
	"tcms/m/internal/automation/interfaces"
)

type Service struct {
	list map[string][]core.Automation
}

// Start launch automation service
func (s *Service) Start(triggerChan chan interfaces.Trigger, errChan chan error) {
	for {
		trigger := <-triggerChan
		err := s.handleTrigger(trigger)
		if err != nil {
			errChan <- err
		}
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

func (s *Service) handleTrigger(trigger interfaces.Trigger) error {
	automationList := s.list[trigger.GetName()]
	if automationList == nil {
		return fmt.Errorf("no automation for trigger %s\n", trigger.GetName())
	}
	for _, automation := range automationList {
		fmt.Printf("Trigger with type %s\n", trigger.GetName())
		err := automation.Execute(trigger)
		if err != nil {
			return err
		}
	}
	return nil
}
