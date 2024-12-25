package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/QBC8-GO-GROUP/GholiBaba/api/service"
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/ticket/domain"

	// "github.com/QBC8-GO-GROUP/GholiBaba/internal/user/domain" //should be imported after having user services
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage/mapper"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage/types"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/logger"
)

func CreateTravel(svcGetter ServiceGetter[*service.TravelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req types.Travel
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.CreateTravel(c.UserContext(), *mapper.TravelStorage2Domain(req))
		if err != nil {
			logger.Error("error in creating role", nil)
			return err
		}
		logger.Info("role created successfully", nil)
		return c.JSON(resp)
	}
}

func GetTravel(svcGetter ServiceGetter[*service.TravelService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		resp, err := svc.GetTravel(c.UserContext(), domain.TravelID(id))
		if err != nil {
			logger.Error("error in fetching role", nil)
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		logger.Info("fetched role successfully", nil)
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
			logger.Error("error in updating role", nil)
			return err
		}
		logger.Info("role updated successfully", nil)
		return nil
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
			logger.Error("error in deleting role", nil)
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		logger.Info("deleted role successfully", nil)
		return nil
	}
}
