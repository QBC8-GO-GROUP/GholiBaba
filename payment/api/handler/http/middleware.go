package http

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/pkg/context"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetUserContext(c *fiber.Ctx) error {
	c.SetUserContext(context.NewAppContext(c.UserContext()))
	return c.Next()
}

func Transaction(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tr := db.Begin()

		context.SetDB(c.UserContext(), tr)

		err := c.Next()

		if err != nil {
			tr.Rollback()
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if err := tr.Commit().Error; err != nil {
			tr.Rollback()
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return nil
	}
}
