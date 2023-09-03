package validation

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/spacenotes-id/SpaceNotes/server/helper/exception"
)

func ValidateBody(c *fiber.Ctx, payload any) error {
	if err := c.BodyParser(payload); err != nil {
		return fiber.
			NewError(fiber.StatusUnprocessableEntity, "Failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	return nil
}
