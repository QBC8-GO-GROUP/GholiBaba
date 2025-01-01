package http

import (
	con "context"

	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/context"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/jwt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func newAuthMiddleware(secret []byte) fiber.Handler {
	return jwtware.New(jwtware.Config{
		AuthScheme:  "Bearer",
		SigningKey:  jwtware.SigningKey{Key: secret},
		Claims:      &jwt.UserClaims{},
		TokenLookup: "header:Authorization",
		SuccessHandler: func(ctx *fiber.Ctx) error {
			userClaims := userClaims(ctx)

			if userClaims == nil {
				return fiber.ErrUnauthorized
			}

			logger := context.GetLogger(ctx.UserContext())
			context.SetLogger(ctx.UserContext(), logger.With("user_id", userClaims.UserID))

			ctx.Locals("user_id", userClaims.UserID)
			ctx.Locals("role", userClaims.Role)
			ctx.SetUserContext(con.WithValue(ctx.UserContext(), "user_id", userClaims.UserID))
			ctx.SetUserContext(con.WithValue(ctx.UserContext(), "role", userClaims.Role))

			// _ = ctx.JSON(fiber.Map{
			// 	"auth_scheme": "Bearer",
			// 	"status":      "authenticated",
			// 	"user": fiber.Map{
			// 		"user_id": userClaims.UserID,
			// 		// "role":    claims.Role,
			// 	},
			// })

			return ctx.Next()
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		},
	})
}
