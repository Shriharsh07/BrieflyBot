package main

import (
	"brieflybot/service"
	ws "brieflybot/websocket"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

func main() {
	// Start WebSocket server
	go func() {
		r := gin.Default()

		r.GET("/ws", func(c *gin.Context) {
			websocket.Handler(ws.HandleWS).ServeHTTP(c.Writer, c.Request)
		})

		r.Run(":8080")
	}()

	// Run your email processor
	service.RunEmailProcessor()

	// 🔒 Block forever (keep app alive)
	select {}
}
