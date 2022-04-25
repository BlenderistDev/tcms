package main

import (
	"context"
	"fmt"
	"time"

	"github.com/BlenderistDev/automation/interfaces"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gopkg.in/errgo.v2/fmt/errors"
	"tcms/internal/automation"
	"tcms/internal/automation/action"
	"tcms/internal/automation/condition"
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

	ctx := context.Background()

	telegram, err := telegramClient.NewTelegram()
	if err != nil {
		panic(errors.Newf("no telegram bridge connection. Error: %v", err))
	}

	connection, err := db.GetConnection(ctx)
	if err != nil {
		panic(errors.Newf("no mongodb connection. Error: %v", err))
	}

	automationRepo := repository.CreateAutomationRepository(connection)
	automations, err := automationRepo.GetAll(ctx)
	if err != nil {
		panic(fmt.Sprintf("automation fetch error. Error: %v", err))
	}

	addConsumer := runKafkaConnection(ctx, log)

	triggerChan := make(chan interfaces.TriggerEvent)
	errChan := make(chan error)

	actionFactory := action.NewFactory(telegram)
	conditionFactory := condition.NewFactory()

	automationService := automation.NewService(telegram, log, errChan, triggerChan, actionFactory, conditionFactory)

	go automationService.RunAutomationService(automations)
	go trigger.StartTelegramUpdateTrigger(addConsumer, triggerChan, log)
	go trigger.StartTimeTrigger(triggerChan, time.Second)
	go func() {
		err := tcms.StartTcmsGrpc(automationRepo, actionFactory)
		if err != nil {
			log.Error(err)
		}
	}()

	select {}
}

func runKafkaConnection(ctx context.Context, log *logrus.Logger) chan chan []uint8 {
	addConsumer := make(chan chan []uint8)
	quitKafka := make(chan bool)
	kafkaError := make(chan error)

	go kafka.CreateKafkaSubscription(ctx, addConsumer, kafkaError, quitKafka)

	go func() {
		for {
			err := <-kafkaError
			log.Error(err)
		}
	}()

	return addConsumer
}
