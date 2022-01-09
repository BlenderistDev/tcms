package telegramClient

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"tcms/m/pkg/telegram"
)

type TelegramClient interface {
	Authorization(phone string) error
	AuthSignIn(code string) error
	GetCurrentUser() (*telegram.User, error)
	Dialogs() (*telegram.DialogsResponse, error)
	SendMessage(peer, message string) error
	MuteUser(id, accessHash string, unMute bool) error
	MuteChat(id string, unMute bool) error
}

type telegramClient struct {
	telegram telegram.TelegramClient
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

	tg := telegram.NewTelegramClient(conn)

	return &telegramClient{telegram: tg}, nil
}

func (t *telegramClient) Authorization(phone string) error {
	_, err := t.telegram.Login(context.Background(), &telegram.LoginMessage{Phone: phone})
	return err
}

func (t *telegramClient) AuthSignIn(code string) error {
	_, err := t.telegram.Sign(context.Background(), &telegram.SignMessage{Code: code})

	return err
}

func (t *telegramClient) GetCurrentUser() (*telegram.User, error) {
	request := telegram.GetUserRequest{Peer: "me"}
	user, err := t.telegram.GetUser(context.Background(), &request)
	return user.GetUser(), err
}

func (t *telegramClient) Dialogs() (*telegram.DialogsResponse, error) {
	dialogs, err := t.telegram.GetDialogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return dialogs, nil
}

func (t *telegramClient) SendMessage(peer, message string) error {
	request := telegram.SendMessageRequest{
		Peer:    peer,
		Message: message,
	}

	_, err := t.telegram.Send(context.Background(), &request)

	return err
}

func (t telegramClient) MuteUser(id, accessHash string, unMute bool) error {
	request := telegram.MuteUserRequest{
		Id:         id,
		AccessHash: accessHash,
		Unmute:     unMute,
	}
	notifySettings, err := t.telegram.MuteUser(context.Background(), &request)

	if err != nil {
		return err
	}

	if !notifySettings.GetSuccess() {
		return fmt.Errorf("error while setting user notify settings")
	}

	return nil
}

func (t telegramClient) MuteChat(id string, unMute bool) error {
	request := telegram.MuteChatRequest{
		Id:     id,
		Unmute: unMute,
	}
	notifySettings, err := t.telegram.MuteChat(context.Background(), &request)

	if err != nil {
		return err
	}

	if !notifySettings.GetSuccess() {
		return fmt.Errorf("error while setting chat notify settings")
	}

	return nil
}
