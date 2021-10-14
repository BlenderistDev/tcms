package telegramClient

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"tcms/m/dry"

	"github.com/xelaj/mtproto/telegram"
)

const (
	//! WARNING: please, DO NOT use this key downloading in production apps, THIS IS ABSOLUTELY INSECURE!
	//! I mean, seriously, this way used just for examples, we can't create most secured app just for
	//! these examples
	publicKeysForExamplesURL = "https://git.io/JtImk"
)

type TelegramClient struct {
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

func NewTelegram() (*TelegramClient, error) {

	wd, err := os.Getwd()
	dry.HandleError(err)

	appId, err := getAppId()
	dry.HandleError(err)

	appHash, err := getAppHash()
	dry.HandleError(err)

	err = prepareStorage()
	if err != nil {
		return nil, err
	}
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

	telegramClient := new(TelegramClient)
	telegramClient.client = client
	telegramClient.appId = appId
	telegramClient.appHash = appHash
	return telegramClient, nil
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
	dry.HandleError(err)

	publicKeys := dir + "/tg_public_keys.pem"

	_, err = os.Stat(publicKeys)
	if err != nil {

		resp, err := http.Get(publicKeysForExamplesURL)
		if err != nil {
			return err
		}
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

func (telegramClient *TelegramClient) Authorization(phone string) (*telegram.AuthSentCode, error) {
	setCode, err := telegramClient.client.AuthSendCode(
		phone, int32(telegramClient.appId), telegramClient.appHash, &telegram.CodeSettings{},
	)
	telegramClient.phone = phone
	dry.HandleError(err)
	return setCode, nil
}

func (telegramClient *TelegramClient) AuthSignIn(code string, sentCode *telegram.AuthSentCode) error {

	auth, err := telegramClient.client.AuthSignIn(
		telegramClient.phone,
		sentCode.PhoneCodeHash,
		code,
	)
	if err == nil {
		fmt.Println(auth)

		fmt.Println("Success! You've signed in!")
	} else {
		fmt.Println(err)
	}
	return err
}

func (telegramClient *TelegramClient) GetUser(username string) (**telegram.ContactsResolvedPeer, error) {
	userData, err := telegramClient.client.ContactsResolveUsername(username)
	return &userData, err
}

func (telegramClient *TelegramClient) GetCurrentUser() (telegram.User, error) {
	fullUser, err := telegramClient.client.UsersGetFullUser(&telegram.InputUserSelf{})
	return fullUser.User, err
}

func (telegramClient *TelegramClient) Contacts() ([]telegram.User, error) {
	resp, err := telegramClient.client.AccountInitTakeoutSession(&telegram.AccountInitTakeoutSessionParams{
		Contacts: true,
	})

	dry.HandleError(err)

	_, err = telegramClient.client.MakeRequest(
		&telegram.InvokeWithTakeoutParams{
			TakeoutID: resp.ID,
			Query:     &telegram.ContactsGetSavedParams{},
		},
	)

	dry.HandleError(err)

	contacts, err := telegramClient.client.ContactsGetContacts(0)

	dry.HandleError(err)

	c := contacts.(*telegram.ContactsContactsObj)

	return c.Users, err
}

func (telegramClient *TelegramClient) Chats() ([]telegram.Chat, error) {

	_, err := telegramClient.client.AccountInitTakeoutSession(&telegram.AccountInitTakeoutSessionParams{
		MessageChats: true,
	})

	dry.HandleError(err)

	var exceptedIds []int32
	chats, err := telegramClient.client.MessagesGetAllChats(exceptedIds)

	dry.HandleError(err)

	c := chats.(*telegram.MessagesChatsObj)

	return c.Chats, err
}

func (telegramClient *TelegramClient) Dialogs() ([]telegram.Dialog, error) {

	_, err := telegramClient.client.AccountInitTakeoutSession(&telegram.AccountInitTakeoutSessionParams{
		MessageChats: true,
	})
	dry.HandleError(err)

	dialogs, err := telegramClient.client.MessagesGetDialogs(&telegram.MessagesGetDialogsParams{
		OffsetPeer: &telegram.InputPeerUser{UserID: 133773580, AccessHash: -315366407886026984},
	})

	dry.HandleError(err)

	c := dialogs.(*telegram.MessagesDialogsSlice)

	return c.Dialogs, err
}

func (telegramClient *TelegramClient) SendMessage(message string, userId int32, accessHash int64) error {

	inputPeerUser := &telegram.InputPeerUser{
		UserID:     userId,
		AccessHash: accessHash,
	}
	messageParams := &telegram.MessagesSendMessageParams{
		Peer:     inputPeerUser,
		Message:  message,
		RandomID: rand.Int63(),
	}

	updates, err := telegramClient.client.MessagesSendMessage(messageParams)
	dry.HandleError(err)
	fmt.Println(updates)
	return err
}
