package main

import (
	"context"
	"github.com/joho/godotenv"
	"tcms/m/internal/automation"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/automation/trigger"
	"tcms/m/internal/db"
	"tcms/m/internal/db/repository"
	"tcms/m/internal/dry"
	"tcms/m/internal/kafka"
	"tcms/m/internal/telegramClient"
	"tcms/m/internal/webserver"
)

func main() {
	// Load values from .env into the system
	err := godotenv.Load()
	dry.HandleErrorPanic(err)

	telegram, err := telegramClient.NewTelegram()
	dry.HandleError(err)

	ctx := context.Background()
	connection, err := db.GetConnection(ctx)
	dry.HandleErrorPanic(err)

	automationRepo := repository.CreateAutomationRepository(connection)

	addConsumer := make(chan chan []uint8)
	quitKafka := make(chan bool)
	kafkaError := make(chan error)

	go kafka.CreateKafkaSubscription(addConsumer, kafkaError, quitKafka)

	triggerChan := make(chan interfaces.Trigger)
	automationService := automation.Service{}

	go automationService.Start(automationRepo, telegram, triggerChan)
	go trigger.StartTelegramTrigger(addConsumer, triggerChan)
	go trigger.StartTimeTrigger(triggerChan)
	go webserver.StartWebServer(telegram, addConsumer)

	select {}
}
