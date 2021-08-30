package main

import (
	"fmt"
	"os"
	"tcms/m/telegram"

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
	telegram, err := telegram.NewTelegram(mtproto1)
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
