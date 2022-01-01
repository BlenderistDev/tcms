package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"tcms/m/internal/dry"
)

// getWcHandler handler for websockets
func getWcHandler(addConsumer chan chan []uint8) func(c *gin.Context) {
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

		ch := make(chan []uint8)
		addConsumer <- ch

		for {
			data := <-ch
			err = ws.WriteMessage(websocket.TextMessage, data)
			dry.HandleError(err)
		}
	}
}
