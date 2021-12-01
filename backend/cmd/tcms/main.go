package main

import (
	"github.com/joho/godotenv"
	"tcms/m/internal/automation"
	"tcms/m/internal/dry"
	"tcms/m/internal/telegramClient"
	"tcms/m/internal/webserver"
)

func main() {
	// Load values from .env into the system
	err := godotenv.Load()
	dry.HandleErrorPanic(err)

	telegram := telegramClient.NewTelegram()
	go telegram.HandleUpdates()
	go automation.UpdateTriggerFactory()
	webserver.StartWebServer(telegram)
}
