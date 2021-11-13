package action

import (
	"fmt"
	"strconv"
	"tcms/m/automation/core"
	"tcms/m/telegramClient"
)

type sendMessageAction struct {
	telegram   *telegramClient.TelegramClient
	peer       string
	message    string
	accessHash string
}

func CreateSendMessageAction() core.Action {
	telegram := telegramClient.NewTelegram()
	return sendMessageAction{
		telegram: telegram,
	}
}

func (a sendMessageAction) Execute(_ core.Trigger) {
	peer, err := strconv.Atoi(a.peer)
	if err != nil {
		fmt.Printf("cannot parse peer %s", a.peer)
	}
	accessHash, err := strconv.Atoi(a.accessHash)
	if err != nil {
		fmt.Printf("cannot parse access hash %s", a.accessHash)
	}
	err = a.telegram.SendMessage(a.message, int32(peer), int64(accessHash))
	if err != nil {
		fmt.Printf("Send message error")
	}
}
