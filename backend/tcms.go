package main

import (
	"fmt"
	"tcms/m/telegram"
	"tcms/m/webserver"
)

func main() {
	telegram, err := telegram.NewTelegram()
	if err != nil {
		fmt.Println(err)
	}

	err = telegram.Connect()
	if err != nil {
		fmt.Println(err)
	}

	webserver.StartWebServer(telegram)

	telegram.CurrentUser()
}
