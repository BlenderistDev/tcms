package automation

import (
	"context"
	"fmt"
	action2 "tcms/m/internal/automation/action"
	condition2 "tcms/m/internal/automation/condition"
	"tcms/m/internal/automation/core"
	"tcms/m/internal/db"
	"tcms/m/internal/db/repository"
	"tcms/m/internal/dry"
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
			action, err := action2.CreateAction(action)
			if err == nil {
				actions[i] = action
			} else {
				fmt.Println(err)
			}
		}

		coreAutomation := core.Automation{Actions: actions}

		if automation.Condition != nil {
			condition, err := condition2.CreateCondition(automation.Condition)
			dry.HandleError(err)
			coreAutomation.Condition = condition
		}
		for _, trigger := range automation.Triggers {
			s.list[trigger] = append(s.list[trigger], coreAutomation)
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
