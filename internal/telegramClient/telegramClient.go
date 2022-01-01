package telegramClient

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	telegram2 "tcms/m/pkg/telegram"
)

type TelegramClient interface {
	Authorization(phone string) error
	AuthSignIn(code string) error
	GetCurrentUser() (*telegram2.User, error)
	Dialogs() (*telegram2.DialogsResponse, error)
	SendMessage(peer, message string) error
}

type telegramClient struct {
	telegram telegram2.TelegramClient
}

func NewTelegram() (TelegramClient, error) {
	host, err := getTelegramBridgeHost()
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	tg := telegram2.NewTelegramClient(conn)

	return &telegramClient{telegram: tg}, nil
}

func (telegramClient *telegramClient) Authorization(phone string) error {
	_, err := telegramClient.telegram.Login(context.Background(), &telegram2.LoginMessage{Phone: phone})
	return err
}

func (telegramClient *telegramClient) AuthSignIn(code string) error {
	_, err := telegramClient.telegram.Sign(context.Background(), &telegram2.SignMessage{Code: code})

	if err == nil {
		fmt.Println("Success! You've signed in!")
	}

	return err
}

func (telegramClient *telegramClient) GetCurrentUser() (*telegram2.User, error) {
	request := telegram2.GetUserRequest{Peer: "me"}
	user, err := telegramClient.telegram.GetUser(context.Background(), &request)
	return user.GetUser(), err
}

func (telegramClient *telegramClient) Dialogs() (*telegram2.DialogsResponse, error) {
	dialogs, err := telegramClient.telegram.GetDialogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return dialogs, nil
}

func (telegramClient *telegramClient) SendMessage(peer, message string) error {
	request := telegram2.SendMessageRequest{
		Peer:    peer,
		Message: message,
	}

	_, err := telegramClient.telegram.Send(context.Background(), &request)

	return err
}
