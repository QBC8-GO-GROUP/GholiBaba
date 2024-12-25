package http

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/nofitication/pkg/broadcast"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

func NotificationSocket(broadcastServer broadcast.Server) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {

		listener := broadcastServer.Subscribe()

		defer broadcastServer.CancelSubscribe(listener)

		var err error
		for {

			m := <-listener
			// need to add some logic

			if err = c.WriteJSON(fiber.Map{
				"message": m,
			}); err != nil {
				log.Println("write:", err)
				break
			}
		}
	})
}
