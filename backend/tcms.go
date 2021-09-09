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

	// err = telegram.Connect()
	// if err != nil {
	// fmt.Println(err)
	// }

	webserver.StartWebServer(telegramClient)

	// telegramClient.CurrentUser()
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"strings"
// 	"time"

// 	"github.com/k0kubun/pp"
// 	"github.com/pkg/errors"
// 	"github.com/xelaj/go-dry"
// 	"github.com/xelaj/mtproto"
// 	"github.com/xelaj/mtproto/telegram"

// 	utils "github.com/xelaj/mtproto/examples/example_utils"
// )

// func main() {

// 	phoneNumber := "79156497607"

// 	// helper variables
// 	appStorage := utils.PrepareAppStorageForExamples()
// 	sessionFile := filepath.Join(appStorage, "session.json")
// 	publicKeys := filepath.Join(appStorage, "tg_public_keys.pem")

// 	// edit these params for you!
// 	client, err := telegram.NewClient(telegram.ClientConfig{
// 		// where to store session configuration. must be set
// 		SessionFile: sessionFile,
// 		// host address of mtproto server. Actually, it can be any mtproxy, not only official
// 		ServerHost: "149.154.167.40:443",
// 		// public keys file is path to file with public keys, which you must get from https://my.telegram.org
// 		PublicKeysFile:  publicKeys,
// 		AppID:           64573,                              // app id, could be find at https://my.telegram.org
// 		AppHash:         "76f28593676e8d6b35a69690882e6ad6", // app hash, could be find at https://my.telegram.org
// 		InitWarnChannel: true,                               // if we want to get errors, otherwise, client.Warnings will be set nil
// 	})

// 	setCode, err := client.AuthSendCode(
// 		phoneNumber, 64573, "76f28593676e8d6b35a69690882e6ad6", &telegram.CodeSettings{},
// 	)

// 	// this part shows how to deal with errors (if you want of course. No one
// 	// like errors, but the can be return sometimes)
// 	if err != nil {
// 		errResponse := &mtproto.ErrResponseCode{}
// 		if !errors.As(err, &errResponse) {
// 			// some strange error, looks like a bug actually
// 			pp.Println(err)
// 			panic(err)
// 		} else {
// 			if errResponse.Message == "AUTH_RESTART" {
// 				println("Oh crap! You accidentally restart authorization process!")
// 				println("You should login only once, if you'll spam 'AuthSendCode' method, you can be")
// 				println("timeouted to loooooooong long time. You warned.")
// 			} else if errResponse.Message == "FLOOD_WAIT_X" {
// 				println("No way... You've reached flood timeout! Did i warn you? Yes, i am. That's what")
// 				println("happens, when you don't listen to me...")
// 				println()
// 				timeoutDuration := time.Second * time.Duration(errResponse.AdditionalInfo.(int))

// 				println("Repeat after " + timeoutDuration.String())
// 			} else {
// 				println("Oh crap! Got strange error:")
// 				pp.Println(errResponse)
// 			}

// 			os.Exit(1)
// 		}
// 	}

// 	fmt.Print("Auth code: ")
// 	code, _ := bufio.NewReader(os.Stdin).ReadString('\n')
// 	code = strings.ReplaceAll(code, "\n", "")

// 	auth, err := client.AuthSignIn(
// 		phoneNumber,
// 		setCode.PhoneCodeHash,
// 		code,
// 	)
// 	if err == nil {
// 		pp.Println(auth)

// 		fmt.Println("Success! You've signed in!")
// 		return
// 	}

// 	// if you don't have password protection — THAT'S ALL! You're already logged in.
// 	// but if you have 2FA, you need to make few more steps:

// 	// could be some errors:
// 	errResponse := &mtproto.ErrResponseCode{}
// 	ok := errors.As(err, &errResponse)
// 	// checking that error type is correct, and error msg is actualy ask for password
// 	if !ok || errResponse.Message != "SESSION_PASSWORD_NEEDED" {
// 		fmt.Println("SignIn failed:", err)
// 		println("\n\nMore info about error:")
// 		pp.Println(err)
// 		os.Exit(1)
// 	}

// 	fmt.Print("Password:")
// 	password, _ := bufio.NewReader(os.Stdin).ReadString('\n')
// 	password = strings.ReplaceAll(password, "\n", "")

// 	accountPassword, err := client.AccountGetPassword()
// 	dry.PanicIfErr(err)

// 	// GetInputCheckPassword is fast response object generator
// 	inputCheck, err := telegram.GetInputCheckPassword(password, accountPassword)
// 	dry.PanicIfErr(err)

// 	auth, err = client.AuthCheckPassword(inputCheck)
// 	dry.PanicIfErr(err)

// 	pp.Println(auth)
// 	fmt.Println("Success! You've signed in!")
// }
