package main

import (
	"fmt"
	"tcms/m/telegramClient"
	"tcms/m/webserver"

	"github.com/joho/godotenv"
)

func main() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	telegram, err := telegramClient.NewTelegram()
	if err != nil {
		fmt.Println(err)
	}
	updateChan := make(chan interface{})
	go telegram.HandleUpdates(updateChan)
	webserver.StartWebServer(telegram, updateChan)
}
