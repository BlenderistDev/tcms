package main

import (
	"fmt"
	"os"
	"tcms/m/telegram"
	"tcms/m/webserver"

	"github.com/shelomentsevd/mtproto"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	// LoadContacts
	mt, err := mtproto.NewMTProto(41994, "269069e15c81241f5670c397941016a2", mtproto.WithAuthFile(wd+"/.telegramgo", false))
	if err != nil {
		fmt.Println(err)
	}
	telegram, err := telegram.NewTelegram(mt)
	if err != nil {
		fmt.Println(err)
	}

	err = telegram.Connect()
	if err != nil {
		fmt.Println(err)
	}

	webserver.StartWebServer(telegram)
}
