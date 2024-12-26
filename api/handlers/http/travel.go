package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/QBC8-GO-GROUP/GholiBaba/api/service"
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/domain"

	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage/mapper"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage/types"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/logger"
)

type PaginationQuery struct {
	Page int `query:"page" default:"1" validate:"gt=0"`
	Size int `query:"size" default:"10" validate:"gt=0"`
}

func CreateTravel(svcGetter ServiceGetter[*service.TravelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req types.Travel
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.CreateTravel(c.UserContext(), *mapper.TravelStorage2Domain(req))
		if err != nil {
			logger.Error("error in creating travel", nil)
			return err
		}
		logger.Info("travel created successfully", nil)
		return c.JSON(resp)
	}
}

func UpdateTravel(svcGetter ServiceGetter[*service.TravelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req types.Travel
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		err := svc.UpdateTravel(c.UserContext(), *mapper.TravelStorage2Domain(req))
		if err != nil {
			logger.Error("error in updating travel", nil)
			return err
		}
		logger.Info("travel updated successfully", nil)
		return nil
	}
}

func GetTravelByID(svcGetter ServiceGetter[*service.TravelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		resp, err := svc.GetTravelByID(c.UserContext(), domain.TravelID(id))
		if err != nil {
			logger.Error("error in fetching travel", nil)
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		logger.Info("fetched travel successfully", nil)
		return c.JSON(resp)
	}
}

func GetTravels(svcGetter ServiceGetter[*service.TravelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		companyId, err := c.ParamsInt("companyId")
		var paginationQuery PaginationQuery
		err = c.QueryParser(&paginationQuery)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		resp, err := svc.GetTravels(c.UserContext(), domain.OwnerID(companyId), paginationQuery.Page, paginationQuery.Size)
		if err != nil {
			logger.Error("error in fetching travel", nil)
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		logger.Info("fetched travel successfully", nil)
		return c.JSON(resp)
	}
}

func DeleteTravel(svcGetter ServiceGetter[*service.TravelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		err = svc.DeleteTravel(c.UserContext(), domain.TravelID(id))
		if err != nil {
			logger.Error("error in deleting travel", nil)
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		logger.Info("deleted travel successfully", nil)
		return nil
	}
}

func BookTravel(svcGetter ServiceGetter[*service.TravelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		err = svc.BookTravel(c.UserContext(), domain.TravelID(id))
		if err != nil {
			logger.Error("error in deleting travel", nil)
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		logger.Info("deleted travel successfully", nil)
		return nil
	}
}

func CancelBooking(svcGetter ServiceGetter[*service.TravelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		err = svc.CancelBooking(c.UserContext(), domain.TravelID(id))
		if err != nil {
			logger.Error("error in deleting travel", nil)
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		logger.Info("deleted travel successfully", nil)
		return nil
	}
}

func ApproveTravel(svcGetter ServiceGetter[*service.TravelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		err = svc.ApproveTravel(c.UserContext(), domain.TravelID(id))
		if err != nil {
			logger.Error("error in deleting travel", nil)
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		logger.Info("deleted travel successfully", nil)
		return nil
	}
}
