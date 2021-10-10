package webserver

import (
	"tcms/m/dry"
	"tcms/m/telegramClient"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xelaj/mtproto/telegram"
)

type loginData struct {
	Phone string `json:"phone" binding:"required"`
}

type signData struct {
	Code string `json:"code" binding:"required"`
}

type sendMessageData struct {
	AccessHash int64  `json:"accessHash" binding:"required"`
	Id         int32  `json:"id" binding:"required"`
	Message    string `json:"message" binding:"required"`
}

func StartWebServer(telegramClient *telegramClient.TelegramClient) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST"},
		AllowHeaders:  []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	var sentCode *telegram.AuthSentCode
	router.POST("/login", func(c *gin.Context) {
		var loginData loginData
		c.BindJSON(&loginData)
		var err error
		sentCode, err = telegramClient.Authorization(loginData.Phone)
		dry.HandleError(err)
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.POST("/sign", func(c *gin.Context) {
		var signData signData
		c.BindJSON(&signData)
		telegramClient.AuthSignIn(signData.Code, sentCode)
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.GET("/me", func(c *gin.Context) {
		user, err := telegramClient.GetCurrentUser()
		dry.HandleError(err)
		c.JSON(200, user)
	})

	router.GET("/contacts", func(c *gin.Context) {
		contacts, err := telegramClient.Contacts()
		dry.HandleError(err)
		c.JSON(200, contacts)
	})

	router.GET("/chats", func(c *gin.Context) {
		chats, err := telegramClient.Chats()
		dry.HandleError(err)
		c.JSON(200, chats)
	})

	router.POST("/message", func(c *gin.Context) {
		var messageData sendMessageData
		c.BindJSON(&messageData)
		telegramClient.SendMessage(messageData.Message, messageData.Id, messageData.AccessHash)
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.Run(":8080")
}
