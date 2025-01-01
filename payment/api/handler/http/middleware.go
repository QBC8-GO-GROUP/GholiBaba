package http

import (
	"errors"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/pkg/context"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetUserContext(c *fiber.Ctx) error {
	c.SetUserContext(context.NewAppContext(c.UserContext()))
	return c.Next()
}

func FindUserId(c *fiber.Ctx) (string, error) {
	headers := c.GetReqHeaders()
	if len(headers) == 0 {
		return "", errors.New("empty header")
	}

	userIds := headers["Userid"]
	if len(userIds) == 0 {
		return "", errors.New("empty userId")
	}

	userId := userIds[0]
	if userId == "" {
		return "", errors.New("no userId")
	}
	return userId, nil
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
