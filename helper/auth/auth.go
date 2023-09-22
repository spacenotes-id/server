package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserIDFromClaims(c *fiber.Ctx) (int, error) {
	user, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Failed to get user id from claims")
	}
	claims, ok2 := user.Claims.(jwt.MapClaims)
	if !ok2 {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Failed to get user id from claims")
	}

	return int(claims["user_id"].(float64)), nil
}
