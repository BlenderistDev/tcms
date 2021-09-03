package webserver

import (
	"fmt"
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

func StartWebServer(telegram *telegram.Telegram) {
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
		var err error
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

	router.GET("/me", func(c *gin.Context) {
		user, err := telegram.CurrentUser()
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(200, user)
	})

	router.GET("/contacts", func(c *gin.Context) {
		contacts, err := telegram.Contacts()
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(200, contacts)
	})

	router.Run(":8080")
}
