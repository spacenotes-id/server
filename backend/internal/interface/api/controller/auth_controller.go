package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/SpaceNotes/backend/internal/application/dto"
	"github.com/tfkhdyt/SpaceNotes/backend/internal/application/usecase"
	"github.com/tfkhdyt/SpaceNotes/backend/pkg/exception"
)

type AuthController struct {
	userUsecase *usecase.UserUsecase `di.inject:"userUsecase"`
}

func (a *AuthController) Register(c *fiber.Ctx) error {
	newUser := new(dto.RegisterRequest)
	if err := c.BodyParser(newUser); err != nil {
		return exception.NewHTTPError(422, "Failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(newUser); err != nil {
		return exception.NewValidationError(err)
	}

	result, err := a.userUsecase.Register(newUser)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(result)
}
