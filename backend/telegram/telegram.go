package telegram

import (
	"fmt"

	"github.com/shelomentsevd/mtproto"
)

type Telegram struct {
	mtproto   *mtproto.MTProto
	state     *mtproto.TL_updates_state
	read      chan struct{}
	stop      chan struct{}
	connected bool
	users     map[int32]mtproto.TL_user
	chats     map[int32]mtproto.TL_chat
	channels  map[int32]mtproto.TL_channel
	phone     string
}

func NewTelegram(pMTProto *mtproto.MTProto) (*Telegram, error) {
	telegram := new(Telegram)
	telegram.mtproto = pMTProto
	telegram.read = make(chan struct{}, 1)
	telegram.stop = make(chan struct{}, 1)
	telegram.users = make(map[int32]mtproto.TL_user)
	telegram.chats = make(map[int32]mtproto.TL_chat)
	telegram.channels = make(map[int32]mtproto.TL_channel)

	return telegram, nil
}

func (telegram *Telegram) Authorization(phone string) (*mtproto.TL_auth_sentCode, error) {
	if phone == "" {
		return nil, fmt.Errorf("Phone number is empty")
	}
	sentCode, err := telegram.mtproto.AuthSendCode(phone)
	if err != nil {
		return nil, err
	}

	if !sentCode.Phone_registered {
		return nil, fmt.Errorf("Phone number isn't registered")
	}
	telegram.phone = phone
	return sentCode, nil
}

func (telegram *Telegram) AuthSignIn(code string, sentCode *mtproto.TL_auth_sentCode) error {
	auth, err := telegram.mtproto.AuthSignIn(telegram.phone, code, sentCode.Phone_code_hash)
	if err != nil {
		return err
	}

	userSelf := auth.User.(mtproto.TL_user)
	telegram.users[userSelf.Id] = userSelf
	message := fmt.Sprintf("Signed in: Id %d name <%s @%s %s>\n", userSelf.Id, userSelf.First_name, userSelf.Username, userSelf.Last_name)
	fmt.Print(message)
	fmt.Println(message)
	fmt.Println(userSelf)
	return nil
}

// Connects to telegram server
func (telegram *Telegram) Connect() error {
	if err := telegram.mtproto.Connect(); err != nil {
		return err
	}
	telegram.connected = true
	fmt.Println("Connected to telegram server")
	return nil
}
