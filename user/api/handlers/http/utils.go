package http

import (
	"fmt"

	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/jwt"

	"github.com/gofiber/fiber/v2"
	jwt2 "github.com/golang-jwt/jwt/v5"
)

// func userClaims(ctx *fiber.Ctx) *jwt.UserClaims {

// 	if u := ctx.Locals("user"); u != nil {
// 		userClaims, ok := u.(*jwt2.Token).Claims.(*jwt.UserClaims)
// 		if ok {
// 			return userClaims
// 		}
// 	}
// 	return nil
// }

func userClaims(ctx *fiber.Ctx) *jwt.UserClaims {
	// Retrieve the token from ctx.Locals("user")
	token, ok := ctx.Locals("user").(*jwt2.Token)
	if !ok || token == nil {
		fmt.Println("Token not found in ctx.Locals or invalid type")
		return nil
	}

	// Extract claims from the token
	claims, ok := token.Claims.(*jwt.UserClaims)
	if !ok {
		fmt.Println("Failed to cast token claims to *jwt.UserClaims")
		return nil
	}

	// Print claims for debugging
	fmt.Printf("Extracted Claims - UserID: %v, Role: %v\n", claims.UserID, claims.Role)

	return claims
}
