package webserver

import (
	"context"
	"github.com/gin-gonic/gin"
	redis2 "github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"net/http"
	"tcms/m/internal/dry"
	"tcms/m/internal/redis"
)

// getWcHandler handler for websockets
func getWcHandler(redisClient redis.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
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
	}
}
