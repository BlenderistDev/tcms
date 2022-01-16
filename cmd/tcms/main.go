package main

import (
	"context"
	"github.com/joho/godotenv"
	action2 "tcms/m/internal/action"
	"tcms/m/internal/automation"
	condition2 "tcms/m/internal/automation/condition"
	"tcms/m/internal/automation/core"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db"
	"tcms/m/internal/db/repository"
	"tcms/m/internal/dry"
	"tcms/m/internal/kafka"
	"tcms/m/internal/telegramClient"
	trigger2 "tcms/m/internal/trigger"
	"tcms/m/internal/webserver"
)

func main() {
	// Load values from .env into the system
	err := godotenv.Load()
	dry.HandleErrorPanic(err)

	telegram, err := telegramClient.NewTelegram()
	dry.HandleError(err)

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

	go func() {
		automationService := automation.Service{}
		automationService.Init()

		for _, auto := range automations {

			coreAutomation := core.Automation{}

			for _, t := range auto.Triggers {
				coreAutomation.AddTrigger(t)
			}

			for _, a := range auto.Actions {
				act, err := action2.CreateAction(a.Name, telegram)
				dry.HandleError(err)
				coreAutomation.AddAction(action2.GetActionWithModel(act, a))
			}

			if auto.Condition != nil {
				condition, err := condition2.CreateCondition(auto.Condition)
				dry.HandleError(err)
				coreAutomation.AddCondition(condition)
			}

			automationService.AddAutomation(coreAutomation)
		}

		automationService.Start(triggerChan)
	}()

	go trigger2.StartTelegramTrigger(addConsumer, triggerChan)
	go trigger2.StartTimeTrigger(triggerChan)
	go webserver.StartWebServer(telegram, addConsumer)

	select {}
}
