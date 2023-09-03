package note

import "github.com/gofiber/fiber/v2"

func GetNoteIDFromParams(c *fiber.Ctx) (int, error) {
	noteID, err := c.ParamsInt("note_id")
	if err != nil {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Failed to get note id")
	}

	return noteID, nil
}
