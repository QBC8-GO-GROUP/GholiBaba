package http

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/cards/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/cards/port"
	walletDomain "github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/domain"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

type CardHandler struct {
	service port.Service
}

func NewCardHandler(service port.Service) *CardHandler {
	return &CardHandler{service: service}
}

func (h *CardHandler) CreateCard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.GetRespHeader("userId", "")
		if userId == "" {
			return fiber.ErrBadRequest
		}

		walId := c.GetRespHeader("walletId", "")
		if walId == "" {
			return fiber.ErrBadRequest
		}

		walletId, err := strconv.Atoi(walId)
		if err != nil {
			return fiber.ErrBadRequest
		}

		var req CreateCardReq
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		if len(req.number) != 16 {
			return fiber.ErrBadRequest
		}

		ctx := c.UserContext()
		err = h.service.CreateCard(ctx, domain.Card{
			CreatedAt: time.Now(),
			Numbers:   req.number,
			WalletId:  walletDomain.WalletID(walletId),
		})
		if err != nil {
			return fiber.ErrInternalServerError
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func (h *CardHandler) ListCard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.GetRespHeader("userId", "")
		if userId == "" {
			return fiber.ErrBadRequest
		}

		ctx := c.UserContext()
		cards, err := h.service.FindCardWithUserId(ctx, userId)
		if err != nil {
			return fiber.ErrInternalServerError
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"cards": cards,
		})
	}
}

func (h *CardHandler) DeleteCard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.GetRespHeader("userId", "")
		if userId == "" {
			return fiber.ErrBadRequest
		}
		ctx := c.UserContext()

		err := h.service.DeleteCardWithUserId(ctx, userId)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusOK)

	}
}

type CreateCardReq struct {
	number string
}

func RegisterCards(api fiber.Router, cardHandler *CardHandler) {
	cards := api.Group("/cards")

	cards.Post("/", cardHandler.CreateCard())
	cards.Get("/", cardHandler.ListCard())
	cards.Delete("/", cardHandler.DeleteCard())
}
