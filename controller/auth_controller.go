package controller

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/spacenotes-id/server/dto"
	"github.com/spacenotes-id/server/helper/exception"
	"github.com/spacenotes-id/server/usecase"
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

func (a *AuthController) getRefreshToken(c *fiber.Ctx) (string, error) {
	payload := new(dto.LogoutRequest)
	if err := c.BodyParser(payload); err != nil {
		rfrsh := c.Cookies("refreshToken")
		if rfrsh == "" {
			return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid refresh token")
		}

		payload.RefreshToken = rfrsh
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return "", exception.NewValidationError(err)
	}

	return payload.RefreshToken, nil
}

func (a *AuthController) Logout(c *fiber.Ctx) error {
	refreshToken, err := a.getRefreshToken(c)
	if err != nil {
		return err
	}

	response, errLogout := a.authUsecase.Logout(refreshToken)
	if errLogout != nil {
		return errLogout
	}

	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    "deleted",
		Path:     "/",
		Expires:  time.Date(2002, time.April, 1, 23, 0, 0, 0, time.UTC),
		HTTPOnly: true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    "deleted",
		Path:     "/",
		Expires:  time.Date(2002, time.April, 1, 23, 0, 0, 0, time.UTC),
		HTTPOnly: true,
	})

	return c.JSON(response)
}

func (a *AuthController) Refresh(c *fiber.Ctx) error {
	refreshToken, err := a.getRefreshToken(c)
	if err != nil {
		return err
	}

	response, errRefresh := a.authUsecase.Refresh(refreshToken)
	if errRefresh != nil {
		return errRefresh
	}

	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    response.Data.AccessToken,
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
	})

	return c.JSON(response)
}
