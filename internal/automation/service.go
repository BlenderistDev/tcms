package automation

import (
	"github.com/BlenderistDev/automation"
	"github.com/BlenderistDev/automation/core"
	"github.com/BlenderistDev/automation/interfaces"
	"github.com/sirupsen/logrus"
	"tcms/internal/automation/action"
	"tcms/internal/automation/condition"
	"tcms/internal/model"
	"tcms/internal/telegramClient"
)

type Service interface {
	RunAutomationService(automations []model.Automation)
}

type service struct {
	tg               telegramClient.TelegramClient
	log              *logrus.Logger
	errChan          chan error
	triggerChan      chan interfaces.TriggerEvent
	actionFactory    action.Factory
	conditionFactory condition.Factory
}

func NewService(
	tg telegramClient.TelegramClient,
	log *logrus.Logger,
	errChan chan error,
	triggerChan chan interfaces.TriggerEvent,
	actionFactory action.Factory,
	conditionFactory condition.Factory,
) Service {
	return &service{
		tg:               tg,
		log:              log,
		errChan:          errChan,
		triggerChan:      triggerChan,
		actionFactory:    actionFactory,
		conditionFactory: conditionFactory,
	}
}

//RunAutomationService launch automation service
func (s *service) RunAutomationService(automations []model.Automation) {

	automationService := automation.Service{}

	for _, automationModel := range automations {
		a, err := s.buildAutomation(&automationModel)
		if err != nil {
			s.log.Error(err)
		}
		automationService.AddAutomation(a)
	}

	go automationService.Start(s.triggerChan, s.errChan)

	go func(errChan chan error) {
		for {
			err := <-errChan
			s.log.Error(err)
		}
	}(s.errChan)
}

func (s *service) buildAutomation(auto *model.Automation) (interfaces.Automation, error) {
	coreAutomation := core.GetAutomation()

	for _, t := range auto.Triggers {
		coreAutomation.AddTrigger(t)
	}

	for _, a := range auto.Actions {
		act, err := s.actionFactory.CreateAction(a.Name)
		if err != nil {
			return nil, err
		}
		coreAutomation.AddAction(action.GetActionWithModel(act, a))
	}

	if auto.Condition != nil && auto.Condition.Name != "" {
		cond, err := s.conditionFactory.CreateCondition(auto.Condition)
		if err != nil {
			return nil, err
		}
		coreAutomation.AddCondition(cond)
	}

	return coreAutomation, nil
}
