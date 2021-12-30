package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"tcms/m/internal/dry"
	"tcms/m/internal/telegramClient"
)

type sendMessageData struct {
	AccessHash int64  `json:"accessHash" binding:"required"`
	Id         int32  `json:"id" binding:"required"`
	Message    string `json:"message" binding:"required"`
}

// getContacts GET /contacts
func getContacts(telegramClient telegramClient.TelegramClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		contacts, err := telegramClient.Contacts()
		dry.HandleError(err)
		c.JSON(200, contacts)
	}
}

// getCurrentUser GET /me
func getCurrentUser(telegramClient telegramClient.TelegramClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		user, err := telegramClient.GetCurrentUser()
		dry.HandleError(err)
		m := jsonpb.Marshaler{
			EmitDefaults: true,
		}
		s, err := m.MarshalToString(user)
		dry.HandleError(err)
		c.JSON(200, s)
	}
}

// sendMessage POST /message
func sendMessage(telegramClient telegramClient.TelegramClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		var messageData sendMessageData
		err := c.BindJSON(&messageData)
		dry.HandleError(err)

		err = telegramClient.SendMessage(string(messageData.Id), messageData.Message)
		dry.HandleError(err)

		c.JSON(200, gin.H{"status": "ok"})
	}
}

// getDialogs GET /dialogs
func getDialogs(telegramClient telegramClient.TelegramClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		dialogs, err := telegramClient.Dialogs()
		dry.HandleError(err)
		m := jsonpb.Marshaler{
			EmitDefaults: true,
		}
		s, err := m.MarshalToString(dialogs)
		dry.HandleError(err)
		c.JSON(200, s)
	}
}

// getChats GET /chats
func getChats(telegramClient telegramClient.TelegramClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		chats, err := telegramClient.Chats()
		dry.HandleError(err)
		c.JSON(200, chats)
	}
}
