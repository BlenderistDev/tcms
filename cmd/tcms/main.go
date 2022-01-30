package main

import (
	"context"
	"github.com/BlenderistDev/automation"
	"github.com/BlenderistDev/automation/core"
	"github.com/BlenderistDev/automation/interfaces"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"tcms/m/internal/automation/action"
	"tcms/m/internal/automation/condition"
	"tcms/m/internal/automation/trigger"
	"tcms/m/internal/connections/db"
	"tcms/m/internal/connections/db/repository"
	"tcms/m/internal/connections/kafka"
	"tcms/m/internal/dry"
	"tcms/m/internal/telegramClient"
	"tcms/m/internal/webserver"
)

func main() {
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

	go func() {
		automationService := automation.Service{}

		for _, auto := range automations {

			coreAutomation := core.GetAutomation()

			for _, t := range auto.Triggers {
				coreAutomation.AddTrigger(t)
			}

			for _, a := range auto.Actions {
				act, err := action.CreateAction(a.Name, telegram)
				if err != nil {
					logrus.Error(err)
				}
				coreAutomation.AddAction(action.GetActionWithModel(act, a))
			}

			if auto.Condition != nil {
				cond, err := condition.CreateCondition(auto.Condition)
				if err != nil {
					logrus.Error(err)
				}
				coreAutomation.AddCondition(cond)
			}

			automationService.AddAutomation(coreAutomation)
		}

		go func(errChan chan error) {
			for {
				err := <-errChan
				logrus.Error(err)
			}
		}(errChan)

		automationService.Start(triggerChan, errChan)

	}()

	go trigger.StartTelegramTrigger(addConsumer, triggerChan)
	go trigger.StartTimeTrigger(triggerChan)
	go webserver.StartWebServer(telegram, addConsumer)

	select {}
}
