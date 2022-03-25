package main

import (
	"context"
	"time"

	"github.com/BlenderistDev/automation"
	"github.com/BlenderistDev/automation/core"
	"github.com/BlenderistDev/automation/interfaces"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"tcms/m/internal/automation/action"
	"tcms/m/internal/automation/condition"
	"tcms/m/internal/automation/trigger"
	"tcms/m/internal/connections/db"
	"tcms/m/internal/connections/kafka"
	"tcms/m/internal/dry"
	"tcms/m/internal/model"
	"tcms/m/internal/repository"
	"tcms/m/internal/tcms"
	"tcms/m/internal/telegramClient"
	"tcms/m/internal/webserver"
)

func main() {
	log := logrus.New()
	// Load values from .env into the system
	err := godotenv.Load()
	dry.HandleErrorPanic(err)

	telegram, err := telegramClient.NewTelegram()
	dry.HandleErrorPanic(err)

	connection, err := db.GetConnection(context.Background())
	dry.HandleErrorPanic(err)

	automationRepo := repository.CreateAutomationRepository(connection)
	automations, err := automationRepo.GetAll(context.Background())
	dry.HandleErrorPanic(err)

	addConsumer := make(chan chan []uint8)
	quitKafka := make(chan bool)
	kafkaError := make(chan error)

	go kafka.CreateKafkaSubscription(addConsumer, kafkaError, quitKafka)

	triggerChan := make(chan interfaces.Trigger)
	errChan := make(chan error)

	go runAutomationService(automations, telegram, log, errChan, triggerChan)
	go trigger.StartTelegramUpdateTrigger(addConsumer, triggerChan, log)
	go trigger.StartTimeTrigger(triggerChan, time.Second)
	go webserver.StartWebServer(telegram, addConsumer)
	go func() {
		err := tcms.StartTcmsGrpc(automationRepo)
		if err != nil {
			log.Error(err)
		}
	}()

	select {}
}

func runAutomationService(automations []model.Automation, telegram telegramClient.TelegramClient, log *logrus.Logger, errChan chan error, triggerChan chan interfaces.Trigger) {

	automationService := automation.Service{}

	for _, auto := range automations {

		coreAutomation := core.GetAutomation()

		for _, t := range auto.Triggers {
			coreAutomation.AddTrigger(t)
		}

		for _, a := range auto.Actions {
			act, err := action.CreateAction(a.Name, telegram)
			if err != nil {
				log.Error(err)
			}
			coreAutomation.AddAction(action.GetActionWithModel(act, a))
		}

		if auto.Condition != nil && auto.Condition.Name != "" {
			cond, err := condition.CreateCondition(auto.Condition)
			if err != nil {
				log.Error(err)
			}
			coreAutomation.AddCondition(cond)
		}

		automationService.AddAutomation(coreAutomation)
	}

	go func(errChan chan error) {
		for {
			err := <-errChan
			log.Error(err)
		}
	}(errChan)

	automationService.Start(triggerChan, errChan)

}
