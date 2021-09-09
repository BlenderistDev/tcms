package telegramClient

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/k0kubun/pp"
	"github.com/xelaj/mtproto/telegram"
)

const (
	//! WARNING: please, DO NOT use this key downloading in production apps, THIS IS ABSOLUTELY INSECURE!
	//! I mean, seriously, this way used just for examples, we can't create most secured app just for
	//! these examples
	publicKeysForExamplesURL = "https://git.io/JtImk"
)

type Telega struct {
	client  *telegram.Client
	phone   string
	appId   int
	appHash string
}

type User struct {
	Id       int32
	Phone    string
	UserName string
}

func NewTelegram() (*Telega, error) {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	appId, err := getAppId()
	if err != nil {
		return nil, err
	}

	appHash, err := getAppHash()
	if err != nil {
		return nil, err
	}
	prepareStorage()
	sessionFile := wd + "/session.json"
	publicKeys := wd + "/tg_public_keys.pem"

	client, _ := telegram.NewClient(telegram.ClientConfig{
		// where to store session configuration. must be set
		SessionFile: sessionFile,
		// host address of mtproto server. Actually, it can be any mtproxy, not only official
		ServerHost: "149.154.167.40:443",
		// public keys file is path to file with public keys, which you must get from https://my.telegram.org
		PublicKeysFile:  publicKeys,
		AppID:           appId,   // app id, could be find at https://my.telegram.org
		AppHash:         appHash, // app hash, could be find at https://my.telegram.org
		InitWarnChannel: true,    // if we want to get errors, otherwise, client.Warnings will be set nil
	})

	telegram := new(Telega)
	telegram.client = client
	telegram.appId = appId
	telegram.appHash = appHash
	return telegram, nil
}

func getAppId() (int, error) {
	appId, exists := os.LookupEnv("TELEGRAM_APP_ID")
	if !exists {
		return 0, fmt.Errorf("no app key")
	}
	return strconv.Atoi(appId)
}

func getAppHash() (string, error) {
	appHash, exists := os.LookupEnv("TELEGRAM_APP_HASH")
	if !exists {
		return "", fmt.Errorf("no app hash")
	}
	return appHash, nil
}

func prepareStorage() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	publicKeys := dir + "/tg_public_keys.pem"

	_, err = os.Stat(publicKeys)
	if err != nil {

		resp, _ := http.Get(publicKeysForExamplesURL)
		defer resp.Body.Close()

		out, _ := os.Create(publicKeys)

		defer out.Close()
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}
	}

	return nil
}

func (telega *Telega) Authorization(phone string) (*telegram.AuthSentCode, error) {
	setCode, err := telega.client.AuthSendCode(
		phone, int32(telega.appId), telega.appHash, &telegram.CodeSettings{},
	)
	telega.phone = phone

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(setCode)
	return setCode, nil
}

func (telega *Telega) AuthSignIn(code string, sentCode *telegram.AuthSentCode) error {

	auth, err := telega.client.AuthSignIn(
		telega.phone,
		sentCode.PhoneCodeHash,
		code,
	)
	if err == nil {
		pp.Println(auth)

		fmt.Println("Success! You've signed in!")
	} else {
		fmt.Println(err)
	}
	return err
}

// // Connects to telegram server
// func (telegram *Telega) Connect() error {
// 	if err := telegram.mtproto.Connect(); err != nil {
// 		return err
// 	}
// 	telegram.connected = true
// 	fmt.Println("Connected to telegram server")
// 	return nil
// }

// func (telegram *Telega) CurrentUser() (*User, error) {
// 	userFull, err := telegram.mtproto.UsersGetFullUsers(mtproto.TL_inputUserSelf{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println(userFull)
// 	user := userFull.User.(mtproto.TL_user)
// 	telegram.users[user.Id] = user

// 	userData := User{Id: user.Id, Phone: user.Phone, UserName: user.Username}
// 	return &userData, nil
// }

// func (telegram *Telega) Contacts() (map[int32]mtproto.TL_user, error) {
// 	tl, err := telegram.mtproto.ContactsGetContacts("")
// 	if err != nil {
// 		return nil, err
// 	}
// 	list, ok := (*tl).(mtproto.TL_contacts_contacts)
// 	if !ok {
// 		return nil, fmt.Errorf("RPC: %#v", tl)
// 	}

// 	contacts := make(map[int32]mtproto.TL_user)
// 	for _, v := range list.Users {
// 		if v, ok := v.(mtproto.TL_user); ok {
// 			contacts[v.Id] = v
// 		}
// 	}

// 	return contacts, nil
// }
