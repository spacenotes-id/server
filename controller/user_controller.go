package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"

	"github.com/spacenotes-id/server/dto"
	"github.com/spacenotes-id/server/helper/auth"
	"github.com/spacenotes-id/server/helper/exception"
	"github.com/spacenotes-id/server/helper/validation"
	"github.com/spacenotes-id/server/usecase"
)

type UserController struct {
	userUsecase *usecase.UserUsecase `di.inject:"userUsecase"`
}

// FindMyAccount godoc
//
//	@Summary		Find my account
//	@Description	Show my account data
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	dto.FindUserByIDResponse
//	@Failure		400	{object}	exception.HttpError
//	@Failure		401	{object}	exception.HttpError
//	@Failure		404	{object}	exception.HttpError
//	@Router			/users/me [get]
//	@Security		ApiKeyAuth
func (u *UserController) FindMyAccount(c *fiber.Ctx) error {
	userID, err := auth.GetUserIDFromClaims(c)
	if err != nil {
		return err
	}

	result, errFind := u.userUsecase.FindUserByID(userID)
	if errFind != nil {
		return errFind
	}

	return c.JSON(result)
}

// UpdateMyAccount godoc
//
//	@Summary		Update my account
//	@Description	Update my account data
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			account	body		dto.UpdateUserRequest	true	"Update user payload"
//	@Success		200		{object}	dto.UpdateUserResponse
//	@Failure		400		{object}	exception.HttpError
//	@Failure		401		{object}	exception.HttpError
//	@Failure		404		{object}	exception.HttpError
//	@Failure		500		{object}	exception.HttpError
//	@Failure		422		{object}	exception.ValErrors
//	@Router			/users/me [put]
//	@Security		ApiKeyAuth
func (u *UserController) UpdateMyAccount(c *fiber.Ctx) error {
	userID, err := auth.GetUserIDFromClaims(c)
	if err != nil {
		return err
	}

	payload := new(dto.UpdateUserRequest)
	if err := validation.ValidateBody(c, payload); err != nil {
		return err
	}

	result, errUpdate := u.userUsecase.UpdateUser(userID, payload)
	if errUpdate != nil {
		return errUpdate
	}

	return c.JSON(result)
}

// UpdateMyEmail godoc
//
//	@Summary		Update my email
//	@Description	Update my account email
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			account	body		dto.UpdateEmailRequest	true	"Update email payload"
//	@Success		200		{object}	dto.UpdateUserResponse
//	@Success		400		{object}	exception.HttpError
//	@Success		404		{object}	exception.HttpError
//	@Success		500		{object}	exception.HttpError
//	@Success		422		{object}	exception.ValErrors
//	@Router			/users/me/email [patch]
//	@Security		ApiKeyAuth
func (u *UserController) UpdateMyEmail(c *fiber.Ctx) error {
	userID, err := auth.GetUserIDFromClaims(c)
	if err != nil {
		return err
	}

	payload := new(dto.UpdateEmailRequest)
	if err := validation.ValidateBody(c, payload); err != nil {
		return err
	}

	result, err := u.userUsecase.UpdateEmail(userID, payload)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (u *UserController) UpdateMyPassword(c *fiber.Ctx) error {
	userID, err := auth.GetUserIDFromClaims(c)
	if err != nil {
		return err
	}

	payload := new(dto.UpdatePasswordRequest)
	if err := c.BodyParser(payload); err != nil {
		return fiber.
			NewError(fiber.StatusUnprocessableEntity, "Failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	result, err := u.userUsecase.UpdatePassword(userID, payload)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (u *UserController) DeleteMyAccount(c *fiber.Ctx) error {
	userID, err := auth.GetUserIDFromClaims(c)
	if err != nil {
		return err
	}

	result, err := u.userUsecase.DeleteUser(userID)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
