package controller

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/SpaceNotes/server/dto"
	"github.com/tfkhdyt/SpaceNotes/server/helper/exception"
	"github.com/tfkhdyt/SpaceNotes/server/usecase"
)

type AuthController struct {
	authUsecase *usecase.AuthUsecase `di.inject:"authUsecase"`
}

func (a *AuthController) Register(c *fiber.Ctx) error {
	newUser := new(dto.RegisterRequest)
	if err := c.BodyParser(newUser); err != nil {
		return fiber.
			NewError(fiber.StatusUnprocessableEntity, "Failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(newUser); err != nil {
		return exception.NewValidationError(err)
	}

	result, err := a.authUsecase.Register(newUser)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(result)
}

func (a *AuthController) Login(c *fiber.Ctx) error {
	payload := new(dto.LoginRequest)
	if err := c.BodyParser(payload); err != nil {
		return fiber.
			NewError(fiber.StatusUnprocessableEntity, "Failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	result, err := a.authUsecase.Login(payload)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    result.Data.AccessToken,
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    result.Data.RefreshToken,
		Path:     "/",
		Expires:  time.Now().Add(720 * time.Hour),
		HTTPOnly: true,
	})

	return c.Status(201).JSON(result)
}
