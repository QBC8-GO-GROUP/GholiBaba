package http

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/nofitication/app"
	"github.com/QBC8-GO-GROUP/GholiBaba/nofitication/pkg/broadcast"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Run(container app.App) {
	router := fiber.New()

	registerAuthAPI(router, container.BroadCast())

	log.Fatal(router.Listen(":8080"))

}

func registerAuthAPI(router fiber.Router, brCast broadcast.Server) {
	router.Use("/notification", UpgradedWebSocket())
	router.Get("/notification/:id", NotificationSocket(brCast))
}
