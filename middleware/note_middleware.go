package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"github.com/spacenotes-id/server/helper/auth"
	"github.com/spacenotes-id/server/usecase"
)

func NoteOwnership(c *fiber.Ctx) error {
	userID := auth.GetUserIDFromClaims(c)
	noteID, err := c.ParamsInt("note_id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to get note id")
	}

	noteUsecase := di.GetInstance("noteUsecase").(*usecase.NoteUsecase)

	if err := noteUsecase.VerifyNoteOwnership(userID, noteID); err != nil {
		return err
	}

	return c.Next()
}
