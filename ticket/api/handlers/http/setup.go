package http

import (
	"fmt"
	"os"
	"time"

	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/app"
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Run(appContainer app.App, config config.ServerConfig) error {
	app := fiber.New(fiber.Config{
		AppName: "GholiBaba v0.0.1",
	})

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} TraceID: ${locals:traceID}\n",
		Output: os.Stdout,
	}))
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        config.RateLimitMaxAttempt,
		Expiration: time.Duration(config.RatelimitTimePeriod) * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			xForwardedFor := c.Get("x-forwarded-for")
			if xForwardedFor == "" {
				return c.IP()
			}
			return xForwardedFor
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendString("STOP SENDING TOO MUCH REQUESTS")
		},
	}))

	api := app.Group("/")

	registerAPI(appContainer, config, api)

	return app.Listen(fmt.Sprintf(":%d", config.HttpPort))
}

func registerAPI(appContainer app.App, cfg config.ServerConfig, api fiber.Router) {
	ticketRouter := api.Group("/ticket")
	ticketSvcGetter := ticketServiceGetter(appContainer, cfg)

	ticketRouter.Post("", BuyTicket(ticketSvcGetter))
	ticketRouter.Get(":id", GetUserTickets(ticketSvcGetter))
	ticketRouter.Delete(":id", CancelTicket(ticketSvcGetter))
}
