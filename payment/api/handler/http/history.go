package http

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history/port"
	"github.com/gofiber/fiber/v2"
)

type HistoryHandler struct {
	service port.Service
}

func NewHistoryHandler(service port.Service) *HistoryHandler {
	return &HistoryHandler{service: service}
}

func (h *HistoryHandler) GetUserHistory() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.GetRespHeader("userId", "")
		if userId == "" {
			return fiber.ErrBadRequest
		}

		ctx := c.UserContext()
		histories, err := h.service.FindHistoryWithUserId(ctx, userId)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if len(histories) == 0 {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"history": histories,
		})

	}
}
