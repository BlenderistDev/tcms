package telegramClient

import (
	"context"
	"fmt"
	"github.com/xelaj/mtproto/telegram"
	"math/rand"
	"os"
	"tcms/m/dry"
	"tcms/m/redis"
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

var client *TelegramClient = nil

func NewTelegram() *TelegramClient {
	if client != nil {
		return client
	}

	wd, err := os.Getwd()
	dry.HandleErrorPanic(err)

	appId, err := getAppId()
	dry.HandleErrorPanic(err)

	appHash, err := getAppHash()
	dry.HandleErrorPanic(err)

	mtprotoHost, err := getMTProtoHost()
	dry.HandleErrorPanic(err)

	prepareStorage()

	sessionFile := wd + "/session.json"
	publicKeys := wd + "/tg_public_keys.pem"

	c, _ := telegram.NewClient(telegram.ClientConfig{
		// where to store session configuration. must be set
		SessionFile: sessionFile,
		// host address of mtproto server. Actually, it can be any mtproxy, not only official
		ServerHost: mtprotoHost,
		// public keys file is path to file with public keys, which you must get from https://my.telegram.org
		PublicKeysFile:  publicKeys,
		AppID:           appId,   // app id, could be find at https://my.telegram.org
		AppHash:         appHash, // app hash, could be find at https://my.telegram.org
		InitWarnChannel: true,    // if we want to get errors, otherwise, client.Warnings will be set nil
	})

	telegramClient := new(TelegramClient)
	telegramClient.client = c
	telegramClient.appId = appId
	telegramClient.appHash = appHash

	client = telegramClient

	return telegramClient
}

func prepareStorage() {
	dir, err := os.Getwd()
	dry.HandleError(err)

	publicKeys := dir + "/tg_public_keys.pem"

	_, err = os.Stat(publicKeys)
	dry.HandleErrorPanic(err)
}

func (telegramClient *TelegramClient) Authorization(phone string) (*telegram.AuthSentCode, error) {
	setCode, err := telegramClient.client.AuthSendCode(
		phone, int32(telegramClient.appId), telegramClient.appHash, &telegram.CodeSettings{},
	)
	telegramClient.phone = phone
	return setCode, err
}

func (telegramClient *TelegramClient) AuthSignIn(code string, sentCode *telegram.AuthSentCode) error {

	_, err := telegramClient.client.AuthSignIn(
		telegramClient.phone,
		sentCode.PhoneCodeHash,
		code,
	)

	if err == nil {
		fmt.Println("Success! You've signed in!")
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

	if err != nil {
		return nil, err
	}

	_, err = telegramClient.client.MakeRequest(
		&telegram.InvokeWithTakeoutParams{
			TakeoutID: resp.ID,
			Query:     &telegram.ContactsGetSavedParams{},
		},
	)

	if err != nil {
		return nil, err
	}

	contacts, err := telegramClient.client.ContactsGetContacts(0)

	if err != nil {
		return nil, err
	}

	c := contacts.(*telegram.ContactsContactsObj)

	return c.Users, err
}

func (telegramClient *TelegramClient) Chats() ([]telegram.Chat, error) {

	_, err := telegramClient.client.AccountInitTakeoutSession(&telegram.AccountInitTakeoutSessionParams{
		MessageChats: true,
	})

	if err != nil {
		return make([]telegram.Chat, 0), err
	}

	var exceptedIds []int32
	chats, err := telegramClient.client.MessagesGetAllChats(exceptedIds)

	if err != nil {
		return make([]telegram.Chat, 0), err
	}

	c := chats.(*telegram.MessagesChatsObj)

	return c.Chats, nil
}

func (telegramClient *TelegramClient) Dialogs() ([]telegram.Dialog, error) {

	_, err := telegramClient.client.AccountInitTakeoutSession(&telegram.AccountInitTakeoutSessionParams{
		MessageChats: true,
	})

	if err != nil {
		return make([]telegram.Dialog, 0), err
	}

	dialogs, err := telegramClient.client.MessagesGetDialogs(&telegram.MessagesGetDialogsParams{
		OffsetPeer: &telegram.InputPeerUser{UserID: 133773580, AccessHash: -315366407886026984},
	})

	if err != nil {
		return make([]telegram.Dialog, 0), err
	}

	c := dialogs.(*telegram.MessagesDialogsSlice)

	return c.Dialogs, nil
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

	_, err := telegramClient.client.MessagesSendMessage(messageParams)

	return err
}

func (telegramClient *TelegramClient) HandleUpdates() {
	var ctx = context.Background()
	redisClient := redis.GetClient()
	telegramClient.client.AddCustomServerRequestHandler(func(i interface{}) bool {
		triggers := recognizeTrigger(i)

		for _, trigger := range triggers {
			_, err := redisClient.Publish(ctx, "update", trigger)
			dry.HandleError(err)
		}

		return false
	})

	// we need to call updates.getState, after that telegram server will send you updates
	_, err := telegramClient.client.UpdatesGetState()
	dry.HandleError(err)
}
