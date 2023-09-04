package space

import "github.com/gofiber/fiber/v2"

func GetSpaceIDFromParams(c *fiber.Ctx) (int, error) {
	spaceID, err := c.ParamsInt("space_id")
	if err != nil {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Failed to get space id")
	}

	return spaceID, nil
}
