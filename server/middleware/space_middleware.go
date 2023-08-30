package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"github.com/tfkhdyt/SpaceNotes/server/helper/auth"
	"github.com/tfkhdyt/SpaceNotes/server/usecase"
)

func SpaceOwnership(c *fiber.Ctx) error {
	userID := auth.GetUserIDFromClaims(c)
	spaceID, err := c.ParamsInt("space_id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to get space id")
	}

	spaceUsecase := di.GetInstance("spaceUsecase").(*usecase.SpaceUsecase)

	if err := spaceUsecase.VerifySpaceOwnership(userID, spaceID); err != nil {
		return err
	}

	return c.Next()
}
