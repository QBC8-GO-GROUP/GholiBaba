package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/api/service"
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/internal/ticket/domain"

	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/pkg/adapters/storage/mapper"
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/pkg/adapters/storage/types"
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/pkg/logger"
)

func BuyTicket(svcGetter ServiceGetter[*service.TicketService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req types.Ticket
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.BuyTicket(c.UserContext(), *&mapper.TicketStorage2Domain(req).TravelID, *&mapper.TicketStorage2Domain(req).UserID)
		if err != nil {
			logger.Error("error in buying ticket", nil)
			return err
		}
		logger.Info("ticket bought successfully", nil)
		return c.JSON(resp)
	}
}

func GetUserTickets(svcGetter ServiceGetter[*service.TicketService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		resp, err := svc.GetUserTickets(c.UserContext(), domain.UserID(id))
		if err != nil {
			logger.Error("error in fetching tickets", nil)
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		logger.Info("fetched user tickets successfully", nil)
		return c.JSON(resp)
	}
}

func CancelTicket(svcGetter ServiceGetter[*service.TicketService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		err = svc.CancelTicket(c.UserContext(), domain.TicketID(id))
		if err != nil {
			logger.Error("error in canceling ticket", nil)
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		logger.Info("canceled ticket successfully", nil)
		return nil
	}
}
