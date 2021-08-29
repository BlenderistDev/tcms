package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shelomentsevd/mtproto"
)

type loginData struct {
	Phone string `json:"phone" binding:"required"`
}

type signData struct {
	Code string `json:"code" binding:"required"`
}

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

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	// LoadContacts
	mtproto1, err := mtproto.NewMTProto(41994, "269069e15c81241f5670c397941016a2", mtproto.WithAuthFile(wd+"/.telegramgo", false))
	if err != nil {
		fmt.Println(err)
	}
	telegram, err := NewTelegram(mtproto1)
	if err != nil {
		fmt.Println(err)
	}

	err = telegram.Connect()
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST"},
		AllowHeaders:  []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	var sentCode *mtproto.TL_auth_sentCode
	router.POST("/login", func(c *gin.Context) {
		var loginData loginData
		c.BindJSON(&loginData)
		sentCode, err = telegram.Authorization(loginData.Phone)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.POST("/sign", func(c *gin.Context) {
		var signData signData
		c.BindJSON(&signData)
		telegram.AuthSignIn(signData.Code, sentCode)
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.Run(":8080")

}
