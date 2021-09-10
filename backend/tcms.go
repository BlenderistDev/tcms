package main

import (
	"fmt"
	"tcms/m/telegramClient"
	"tcms/m/webserver"

	"github.com/joho/godotenv"
)

func main() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	telegramClient, err := telegramClient.NewTelegram()
	if err != nil {
		fmt.Println(err)
	}
	webserver.StartWebServer(telegramClient)
}
