package controller

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"

	"github.com/spacenotes-id/server/dto"
	"github.com/spacenotes-id/server/helper/exception"
	"github.com/spacenotes-id/server/helper/validation"
	"github.com/spacenotes-id/server/usecase"
)

type AuthController struct {
	authUsecase *usecase.AuthUsecase `di.inject:"authUsecase"`
}

// Register godoc
//
//	@Summary		Register
//	@Description	Register a new account
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			account	body		dto.RegisterRequest	true	"Request body"
//	@Success		201		{object}	dto.RegisterResponse
//	@Failure		422		{object}	exception.ValErrors
//	@Failure		400		{object}	exception.HttpError
//	@Failure		500		{object}	exception.HttpError
//	@Router			/auth/register [post]
func (a *AuthController) Register(c *fiber.Ctx) error {
	newUser := new(dto.RegisterRequest)
	if err := validation.ValidateBody(c, newUser); err != nil {
		return err
	}

	result, err := a.authUsecase.Register(newUser)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(result)
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login to get access token and refresh token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		dto.LoginRequest	true	"Credentials"
//	@Success		201			{object}	dto.LoginResponse
//	@Failure		422			{object}	exception.ValErrors
//	@Failure		400			{object}	exception.HttpError
//	@Failure		404			{object}	exception.HttpError
//	@Failure		500			{object}	exception.HttpError
//	@Router			/auth/login [post]
func (a *AuthController) Login(c *fiber.Ctx) error {
	payload := new(dto.LoginRequest)
	if err := validation.ValidateBody(c, payload); err != nil {
		return err
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
		SameSite: "None",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    result.Data.RefreshToken,
		Path:     "/",
		Expires:  time.Now().Add(720 * time.Hour),
		HTTPOnly: true,
		SameSite: "None",
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

// Logout godoc
//
//	@Summary		Logout
//	@Description	Remove refresh token from database
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			refreshToken	body		dto.LogoutRequest	true	"Refresh token"
//	@Success		200				{object}	dto.LogoutResponse
//	@Failure		422				{object}	exception.ValErrors
//	@Failure		401				{object}	exception.HttpError
//	@Failure		404				{object}	exception.HttpError
//	@Failure		500				{object}	exception.HttpError
//	@Router			/auth/logout [delete]
//	@Security		ApiKeyAuth
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
		SameSite: "None",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    "deleted",
		Path:     "/",
		Expires:  time.Date(2002, time.April, 1, 23, 0, 0, 0, time.UTC),
		HTTPOnly: true,
		SameSite: "None",
	})

	return c.JSON(response)
}

// Refresh godoc
//
//	@Summary		Refresh
//	@Description	Refresh access token using refresh token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			refreshToken	body		dto.RefreshRequest	true	"Refresh token"
//	@Success		200				{object}	dto.RefreshResponse
//	@Failure		422				{object}	exception.ValErrors
//	@Failure		401				{object}	exception.HttpError
//	@Failure		400				{object}	exception.HttpError
//	@Failure		404				{object}	exception.HttpError
//	@Failure		500				{object}	exception.HttpError
//	@Router			/auth/refresh [patch]
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
		SameSite: "None",
	})

	return c.JSON(response)
}
