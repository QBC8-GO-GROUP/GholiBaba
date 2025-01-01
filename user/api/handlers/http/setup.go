package http

import (
	"fmt"

	"github.com/QBC8-GO-GROUP/GholiBaba/api/service"
	"github.com/QBC8-GO-GROUP/GholiBaba/app"
	"github.com/QBC8-GO-GROUP/GholiBaba/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	router := fiber.New()

	api := router.Group("/api/v1")

	api.Post("/sign-up", SignUp(service.NewUserService(appContainer.UserService(),
		cfg.Secret, cfg.AuthExpMinute, cfg.AuthRefreshMinute)))

	api.Post("/sign-in", SignIn(service.NewUserService(appContainer.UserService(),
		cfg.Secret, cfg.AuthExpMinute, cfg.AuthRefreshMinute)))

	// app.Patch("/:id/role", newAuthMiddleware([]byte(cfg.Secret)),
	// 	UpdateUserRoleHandler(service.NewUserService(appContainer.UserService(),
	// 		cfg.Secret, cfg.AuthExpMinute, cfg.AuthRefreshMinute)))

	api.Get("/test", newAuthMiddleware([]byte(cfg.Secret)), TestHandler)

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}
