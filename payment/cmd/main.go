package main

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/api/handler/http"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/app"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/config"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/cards"
	cardsDomain "github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/cards/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history"
	historyDomain "github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet"
	walletDomain "github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/domain"
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

	fiberApp := fiber.New()

	fiberApp.Use(logger.New())
	fiberApp.Use(recover2.New())

	err := application.DB().AutoMigrate(
		&cardsDomain.Card{},
		&historyDomain.History{},
		&walletDomain.Wallet{},
	)

	if err != nil {
		log.Fatalf("Error migrating schema: %v", err)
	}

	// Log successful migration
	log.Println("Database migrated successfully!")

	http.RegisterHistory(fiberApp, historyHandler)
	http.RegisterCards(fiberApp, cardHandler)

	transaction := http.Transaction(application.DB())
	http.RegisterWaller(fiberApp, transaction, walletHandler)

	log.Fatal(fiberApp.Listen("8080"))

}
