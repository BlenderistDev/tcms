package main

import (
	"github.com/joho/godotenv"
	"tcms/m/internal/automation"
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

	addConsumer := make(chan chan []uint8)
	quitKafka := make(chan bool)
	kafkaError := make(chan error)

	go kafka.CreateKafkaSubscription(addConsumer, kafkaError, quitKafka)
	go automation.UpdateTriggerFactory(addConsumer)
	go webserver.StartWebServer(telegram, addConsumer)

	select {}
}
