package http

import (
	con "context"
	"errors"
	"fmt"
	"time"

	"github.com/QBC8-GO-GROUP/GholiBaba/api/pb"
	"github.com/QBC8-GO-GROUP/GholiBaba/api/service"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/context"
	"github.com/gofiber/fiber/v2"
)

func SignUp(svc *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var req pb.UserSignUpRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.SignUp(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, service.ErrUserCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)

	}
}

func SignIn(svc *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var req pb.UserSignInRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.SignIn(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, service.ErrUserNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			if errors.Is(err, service.ErrInvalidUserPassword) {
				return fiber.NewError(fiber.StatusUnauthorized, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func UpdateUserRoleHandler(svc *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			Role string `json:"role"`
		}

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		roleEnum := pb.Role_value[req.Role]

		changeRoleReq := &pb.ChangeRoleRequest{
			Role: pb.Role(roleEnum),
		}

		ctx := con.WithValue(c.UserContext(), "user_id", c.Locals("user_id"))
		// ctx = con.WithValue(ctx, "role", c.Locals("role"))

		err := svc.UpdateUserRoleHandler(ctx, changeRoleReq)

		if err != nil {
			if errors.Is(err, service.ErrUserNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "User not found",
				})
			}
			if errors.Is(err, service.ErrInvalidRole) {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid role provided",
				})
			}
		}

		return c.JSON(fiber.Map{
			"message": "User role updated successfully",
		})
	}
}

func TestHandler(ctx *fiber.Ctx) error {
	logger := context.GetLogger(ctx.UserContext())

	fmt.Println("ctx.UserContext():", ctx.UserContext())
	fmt.Println("userIDFromContext:", ctx.UserContext().Value("user_id"))
	fmt.Println("userRoleFromContext:", ctx.UserContext().Value("role"))

	logger.Info("User validation process ... ",
		"authHeader", ctx.Get("Authorization")[:6],
		"userAgent", ctx.Get("User-Agent")[:14],
		"userID", ctx.Locals("user_id"),
		"role", ctx.Locals("role"),
		"time", time.Now().Format(time.DateTime))

	return nil
}
