package main

import (
	"context"
	"time"

	"github.com/BlenderistDev/automation/interfaces"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"tcms/m/internal/automation"
	"tcms/m/internal/automation/trigger"
	"tcms/m/internal/connections/db"
	"tcms/m/internal/connections/kafka"
	"tcms/m/internal/dry"
	"tcms/m/internal/repository"
	"tcms/m/internal/tcms"
	"tcms/m/internal/telegramClient"
)

func main() {
	log := logrus.New()
	// Load values from .env into the system
	err := godotenv.Load()
	if err != nil {
		log.Error(err)
	}

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

	triggerChan := make(chan interfaces.TriggerEvent)
	errChan := make(chan error)

	go automation.RunAutomationService(automations, telegram, log, errChan, triggerChan)
	go trigger.StartTelegramUpdateTrigger(addConsumer, triggerChan, log)
	go trigger.StartTimeTrigger(triggerChan, time.Second)
	go func() {
		err := tcms.StartTcmsGrpc(automationRepo)
		if err != nil {
			log.Error(err)
		}
	}()

	select {}
}
