package telegram

import (
	"fmt"
	"github.com/shelomentsevd/mtproto"
	"os"
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

type User struct {
	Id       int32
	Phone    string
	UserName string
}

func NewTelegram() (*Telegram, error) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	// LoadContacts
	mt, err := mtproto.NewMTProto(41994, "269069e15c81241f5670c397941016a2", mtproto.WithAuthFile(wd+"/.telegramgo", false))
	if err != nil {
		fmt.Println(err)
	}
	telegram := new(Telegram)
	telegram.mtproto = mt
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

func (telegram *Telegram) CurrentUser() (*User, error) {
	userFull, err := telegram.mtproto.UsersGetFullUsers(mtproto.TL_inputUserSelf{})
	if err != nil {
		return nil, err
	}
	fmt.Println(userFull)
	user := userFull.User.(mtproto.TL_user)
	telegram.users[user.Id] = user

	userData := User{Id: user.Id, Phone: user.Phone, UserName: user.Username}
	return &userData, nil
}
