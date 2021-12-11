package webserver

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	redis2 "github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/xelaj/mtproto/telegram"
	"net/http"
	"tcms/m/internal/dry"
	"tcms/m/internal/redis"
	"tcms/m/internal/telegramClient"
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

func StartWebServer(telegramClient telegramClient.TelegramClient, redisClient redis.Client) {
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
		err := c.BindJSON(&loginData)
		dry.HandleError(err)

		sentCode, err = telegramClient.Authorization(loginData.Phone)
		dry.HandleError(err)
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.POST("/sign", func(c *gin.Context) {
		var signData signData
		err := c.BindJSON(&signData)
		dry.HandleError(err)

		err = telegramClient.AuthSignIn(signData.Code, sentCode)
		dry.HandleError(err)

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

	router.GET("/dialogs", func(c *gin.Context) {
		dialogs, err := telegramClient.Dialogs()
		dry.HandleError(err)
		c.JSON(200, dialogs)
	})

	router.POST("/message", func(c *gin.Context) {
		var messageData sendMessageData
		err := c.BindJSON(&messageData)
		dry.HandleError(err)

		err = telegramClient.SendMessage(messageData.Message, messageData.Id, messageData.AccessHash)
		dry.HandleError(err)

		c.JSON(200, gin.H{"status": "ok"})
	})

	router.GET("/ws", func(c *gin.Context) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		dry.HandleError(err)

		defer func(ws *websocket.Conn) {
			err := ws.Close()
			dry.HandleError(err)
		}(ws)

		var ctx = context.Background()

		pubsub := redisClient.Subscribe(ctx, "update")
		defer func(pubsub *redis2.PubSub) {
			err := pubsub.Close()
			dry.HandleError(err)
		}(pubsub)

		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			dry.HandleError(err)
			err = ws.WriteJSON(msg.Payload)
			dry.HandleError(err)
		}
	})

	host, err := getApiHost()
	dry.HandleErrorPanic(err)

	dry.HandleErrorPanic(router.Run(host))
}
