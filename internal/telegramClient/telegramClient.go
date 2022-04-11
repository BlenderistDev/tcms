package telegramClient

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"tcms/pkg/telegram"
)

type TelegramClient interface {
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
	res, err := t.telegram.MuteUser(context.Background(), &request)

	if err != nil {
		return err
	}

	if !res.GetSuccess() {
		return fmt.Errorf("error while setting user notify settings")
	}

	return nil
}

func (t telegramClient) MuteChat(id string, unMute bool) error {
	request := telegram.MuteChatRequest{
		Id:     id,
		Unmute: unMute,
	}
	res, err := t.telegram.MuteChat(context.Background(), &request)

	if err != nil {
		return err
	}

	if !res.GetSuccess() {
		return fmt.Errorf("error while setting chat notify settings")
	}

	return nil
}
