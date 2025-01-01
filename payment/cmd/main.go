package main

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/api/handler/http"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/app"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/config"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/cards"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/pkg/adapters/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func main() {
	cfg := config.MustReadConfig("config.json")
	application := app.MustNewApp(cfg)

	walletRepo := storage.NewWalletRepo(application.DB())
	cardRepo := storage.NewCardsRepo(application.DB())
	historyRepo := storage.NewHistoryRepo(application.DB())

	walletService := wallet.NewService(walletRepo)
	cardService := cards.NewService(cardRepo)
	historyService := history.NewService(historyRepo)

	walletHandler := http.NewWalletHandler(walletService)
	cardHandler := http.NewCardHandler(cardService)
	historyHandler := http.NewHistoryHandler(historyService)

	fiberApp := fiber.New(fiber.Config{
		DisableStartupMessage: false,
	})
	router := fiberApp.Group("/")

	fiberApp.Use(recover2.New())
	fiberApp.Use(logger.New(logger.Config{}))

	fiberApp.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "service is up and running",
		})
	})

	transaction := http.Transaction(application.DB())
	http.RegisterWaller(router, transaction, walletHandler)

	http.RegisterHistory(router, historyHandler)
	http.RegisterCards(router, cardHandler)

	log.Fatal(fiberApp.Listen(":8080"))

}
