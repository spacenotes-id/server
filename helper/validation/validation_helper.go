package validation

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"

	"github.com/spacenotes-id/server/database/postgres/sqlc"
	"github.com/spacenotes-id/server/helper/exception"
)

func ValidateBody[T any](c *fiber.Ctx, payload T) error {
	if err := c.BodyParser(payload); err != nil {
		return fiber.
			NewError(fiber.StatusBadRequest, "Failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	return nil
}

func ValidateStatusQuery(status sqlc.Status) error {
	if ok := govalidator.IsIn(
		string(status),
		"normal", "favorite", "archived", "trashed", "",
	); !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid status")
	}

	return nil
}
