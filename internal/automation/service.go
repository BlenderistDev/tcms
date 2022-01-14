package automation

import (
	"context"
	"fmt"
	condition2 "tcms/m/internal/automation/condition"
	"tcms/m/internal/automation/core"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db/repository"
	"tcms/m/internal/dry"
	"tcms/m/internal/telegramClient"
)

type Service struct {
	actionMap map[string]interfaces.Action
	list      map[string][]core.Automation
}

// Start launch automation service
func (s *Service) Start(automationRepo repository.AutomationRepository, telegram telegramClient.TelegramClient, triggerChan chan interfaces.Trigger) {
	automations, err := automationRepo.GetAll(context.Background())
	dry.HandleErrorPanic(err)

	s.list = make(map[string][]core.Automation, len(automations))

	for _, automation := range automations {
		actions := make([]interfaces.ActionWithModel, len(automation.Actions))
		for i, action := range automation.Actions {
			if err == nil {
				actions[i] = core.GetActionWithModel(s.getAction(action.Name), action)
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

	for {
		trigger := <-triggerChan
		s.HandleTrigger(trigger)
	}
}

func (s *Service) AddAction(name string, action interfaces.Action) {
	if s.actionMap == nil {
		s.actionMap = map[string]interfaces.Action{}
	}
	s.actionMap[name] = action
}

func (s *Service) getAction(name string) interfaces.Action {
	return s.actionMap[name]
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
