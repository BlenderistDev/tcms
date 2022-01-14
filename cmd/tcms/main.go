package main

import (
	"context"
	"github.com/joho/godotenv"
	"tcms/m/internal/automation"
	"tcms/m/internal/automation/action"
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

	connection, err := db.GetConnection(context.Background())
	dry.HandleErrorPanic(err)

	automationRepo := repository.CreateAutomationRepository(connection)

	addConsumer := make(chan chan []uint8)
	quitKafka := make(chan bool)
	kafkaError := make(chan error)

	go kafka.CreateKafkaSubscription(addConsumer, kafkaError, quitKafka)

	triggerChan := make(chan interfaces.Trigger)

	go func() {
		automationService := automation.Service{}
		automationService.AddAction("sendMessage", action.CreateSendMessageAction(telegram))
		automationService.AddAction("muteUser", action.CreateMuteUserAction(telegram))
		automationService.AddAction("muteChat", action.CreateMuteChatAction(telegram))
		automationService.Start(automationRepo, triggerChan)
	}()

	go trigger.StartTelegramTrigger(addConsumer, triggerChan)
	go trigger.StartTimeTrigger(triggerChan)
	go webserver.StartWebServer(telegram, addConsumer)

	select {}
}
