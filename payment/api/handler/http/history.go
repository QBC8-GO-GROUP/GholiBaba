package http

import (
	historyDomain "github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history/port"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func (h *HistoryHandler) ApproveFactor() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.GetRespHeader("role", "")
		if userRole != "admin" {
			return fiber.NewError(fiber.StatusForbidden, "You do not have permission to approve transactions")
		}

		historyIdStr := c.Params("id")
		if historyIdStr == "" {
			return fiber.NewError(fiber.StatusBadRequest, "History ID is required")
		}

		historyId, err := uuid.Parse(historyIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid history ID")
		}

		histories, err := h.service.FindHistoryWithId(c.UserContext(), historyDomain.HistoryId(historyId))
		if err != nil || len(histories) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "History record not found",
			})
		}

		// We assume there's only one history record with the given ID
		historyRecord := histories[0]

		// Check if the history is already approved
		if historyRecord.IsApproved {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "This transaction is already approved",
			})
		}

		historyRecord.IsApproved = true

		err = h.service.UpdateHistory(c.UserContext(), historyRecord)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to approve the history")
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Transaction approved successfully",
			"history": historyRecord,
		})
	}
}

func RegisterHistory(api fiber.Router, historyHandler *HistoryHandler) {
	historyApi := api.Group("/history")
	historyApi.Put("/:id/approve", historyHandler.ApproveFactor())
}
