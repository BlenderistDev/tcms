package main

import (
	"context"
	"fmt"
	"time"

	"github.com/BlenderistDev/automation/interfaces"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"tcms/internal/automation"
	"tcms/internal/automation/trigger"
	"tcms/internal/connections/db"
	"tcms/internal/connections/kafka"
	"tcms/internal/repository"
	"tcms/internal/tcms"
	"tcms/internal/telegramClient"
)

func main() {
	log := logrus.New()
	// Load values from .env into the system
	err := godotenv.Load()
	if err != nil {
		log.Error(err)
	}

	telegram, err := telegramClient.NewTelegram()
	if err != nil {
		panic(fmt.Sprintf("no telegram bridge connection. Error: %v", err))
	}

	connection, err := db.GetConnection(context.Background())
	if err != nil {
		panic(fmt.Sprintf("no mongodb connection. Error: %v", err))
	}

	automationRepo := repository.CreateAutomationRepository(connection)
	automations, err := automationRepo.GetAll(context.Background())
	if err != nil {
		panic(fmt.Sprintf("automation fetch error. Error: %v", err))
	}

	addConsumer := runKafkaConnection(log)

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

func runKafkaConnection(log *logrus.Logger) chan chan []uint8 {
	addConsumer := make(chan chan []uint8)
	quitKafka := make(chan bool)
	kafkaError := make(chan error)

	go kafka.CreateKafkaSubscription(addConsumer, kafkaError, quitKafka)

	go func() {
		for {
			err := <-kafkaError
			log.Error(err)
		}
	}()

	return addConsumer
}
