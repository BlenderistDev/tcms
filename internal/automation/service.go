package automation

import (
	"github.com/BlenderistDev/automation"
	"github.com/BlenderistDev/automation/core"
	"github.com/BlenderistDev/automation/interfaces"
	"github.com/sirupsen/logrus"
	"tcms/m/internal/automation/action"
	"tcms/m/internal/automation/condition"
	"tcms/m/internal/model"
	"tcms/m/internal/telegramClient"
)

//RunAutomationService launch automation service
func RunAutomationService(automations []model.Automation, telegram telegramClient.TelegramClient, log *logrus.Logger, errChan chan error, triggerChan chan interfaces.TriggerEvent) {

	automationService := automation.Service{}

	for _, automationModel := range automations {
		a, err := buildAutomation(&automationModel, telegram)
		if err != nil {
			log.Error(err)
		}
		automationService.AddAutomation(a)
	}

	go automationService.Start(triggerChan, errChan)

	go func(errChan chan error) {
		for {
			err := <-errChan
			log.Error(err)
		}
	}(errChan)
}

func buildAutomation(auto *model.Automation, telegram telegramClient.TelegramClient) (interfaces.Automation, error) {
	coreAutomation := core.GetAutomation()

	for _, t := range auto.Triggers {
		coreAutomation.AddTrigger(t)
	}

	for _, a := range auto.Actions {
		act, err := action.CreateAction(a.Name, telegram)
		if err != nil {
			return nil, err
		}
		coreAutomation.AddAction(action.GetActionWithModel(act, a))
	}

	if auto.Condition != nil && auto.Condition.Name != "" {
		cond, err := condition.CreateCondition(auto.Condition)
		if err != nil {
			return nil, err
		}
		coreAutomation.AddCondition(cond)
	}

	return coreAutomation, nil
}
