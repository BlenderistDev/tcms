package automation

import (
	"context"
	"fmt"
	action2 "tcms/m/automation/action"
	"tcms/m/automation/core"
	"tcms/m/db"
	"tcms/m/db/repository"
	"tcms/m/dry"
)

type Service struct {
	list map[string][]core.Automation
}

// Start launch automation service
func (s *Service) Start() {
	ctx := context.Background()
	connection := db.GetConnection(ctx)

	automationRepo := repository.CreateAutomationRepository(connection)
	automations, err := automationRepo.GetAll(ctx)
	if err != nil {
		dry.HandleErrorPanic(err)
	}

	s.list = make(map[string][]core.Automation, len(automations))

	for _, automation := range automations {
		actions := make([]core.Action, len(automation.Actions))
		for i, action := range automation.Actions {
			action, err := action2.CreateAction(action.Name)
			if err == nil {
				actions[i] = action
			} else {
				fmt.Println(err)
			}
		}
		for _, trigger := range automation.Triggers {
			s.list[trigger] = append(s.list[trigger], core.Automation{Actions: actions})
		}
	}
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
