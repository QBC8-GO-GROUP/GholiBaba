package http

import (
	"fmt"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history"
	historyDomain "github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/port"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/pkg/adapters/storage"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/pkg/context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"time"
)

type WalletHandler struct {
	walletService port.Service
}

func NewWalletHandler(walletService port.Service) *WalletHandler {
	return &WalletHandler{
		walletService: walletService,
	}
}

func Pay() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.GetRespHeader("userId", "")
		if userId == "" {
			return fiber.ErrBadRequest
		}

		var req PayReq
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		db := context.GetDB(c.UserContext())
		if db == nil {
			return fiber.ErrInternalServerError
		}

		walletRepo := storage.NewWalletRepo(db)
		walletService := wallet.NewService(walletRepo)

		historyRepo := storage.NewHistoryRepo(db)
		historyService := history.NewService(historyRepo)

		sourceWallet, err := walletService.FindUserWallet(c.UserContext(), userId)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "source wallet not found",
			})
		}

		destinationWallet, err := walletService.FindUserWallet(c.UserContext(), req.Destination)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "destination wallet not found",
			})
		}

		if req.Price > sourceWallet.Money {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "source wallet amount is lower than destination wallet amount",
			})
		}

		sourceWallet.Money = sourceWallet.Money - req.Price
		destinationWallet.Money = destinationWallet.Money + req.Price

		err = walletService.UpdateWallet(c.UserContext(), sourceWallet)
		if err != nil {
			return fiber.ErrInternalServerError
		}

		err = walletService.UpdateWallet(c.UserContext(), destinationWallet)
		if err != nil {
			return fiber.ErrInternalServerError
		}

		uid, err := uuid.NewUUID()
		if err != nil {
			return fiber.ErrInternalServerError
		}
		createHistory, err := historyService.CreateHistory(c.UserContext(), historyDomain.History{
			CreatedAt:   time.Now(),
			Code:        uid,
			IsApproved:  false,
			Price:       req.Price,
			Source:      sourceWallet.Id,
			Destination: destinationWallet.Id,
			Title:       req.Title,
			Description: req.Description,
		})

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		_ = createHistory

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code": uid.String(),
		})
	}
}

func AddMoney() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId, err := FindUserId(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "user id not found",
			})
		}

		var req AddMoneyReq
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "bad body request",
			})
		}

		db := context.GetDB(c.UserContext())

		walletRepo := storage.NewWalletRepo(db)
		walletService := wallet.NewService(walletRepo)

		historyRepo := storage.NewHistoryRepo(db)
		historyService := history.NewService(historyRepo)

		userWallet, err := walletService.FindUserWallet(c.UserContext(), userId)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "user wallet not found",
			})
		}

		userWallet.Money = userWallet.Money + req.Money

		err = walletService.UpdateWallet(c.UserContext(), userWallet)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "user wallet can not be updated",
			})
		}

		code, err := uuid.NewUUID()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		createHistory, err := historyService.CreateHistory(c.UserContext(), historyDomain.History{
			CreatedAt:   time.Now(),
			Code:        code,
			IsApproved:  true,
			Price:       req.Money,
			Source:      0,
			Destination: userWallet.Id,
			Title:       fmt.Sprintf("add money to %v", userWallet.UserId),
			Description: fmt.Sprintf("add money %v to %v", req.Money, userWallet.UserId),
		})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "can not save history",
			})
		}

		_ = createHistory

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code": code.String(),
		})
	}
}

func CreateWallet(h *WalletHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId, err := FindUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "user wallet not found",
			})
		}

		code, err := uuid.NewUUID()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "can create wallet id",
			})
		}

		var newWallet = domain.Wallet{
			UserId:    userId,
			CreatedAt: time.Now(),
			Code:      code,
			Money:     0,
		}
		ctx := c.UserContext()
		err = h.walletService.CreateWallet(ctx, newWallet)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "db problem",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"code": code,
		})
	}
}

func RegisterWaller(api fiber.Router, transaction fiber.Handler, walletHandler *WalletHandler) {
	group := api.Group("/wallet", SetUserContext)

	//group.Post("/", func(ctx *fiber.Ctx) error {
	//	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
	//		"success": true,
	//	})
	//})
	group.Post("/", CreateWallet(walletHandler))

	group.Use(transaction)

	group.Post("money", AddMoney())
	group.Post("pay", Pay())
}

type AddMoneyReq struct {
	Money float64 `json:"money"`
}

type PayReq struct {
	Price           float64 `json:"price"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	Destination     string  `json:"destination"`
	DestinationName string  `json:"destinationName"`
}
