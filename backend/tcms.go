package main

import (
	"tcms/m/automation"
	"tcms/m/dry"
	"tcms/m/telegramClient"
	"tcms/m/webserver"

	"github.com/joho/godotenv"
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
