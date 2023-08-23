package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/SpaceNotes/server/internal/application/dto"
	"github.com/tfkhdyt/SpaceNotes/server/internal/application/usecase"
	"github.com/tfkhdyt/SpaceNotes/server/pkg/auth"
	"github.com/tfkhdyt/SpaceNotes/server/pkg/exception"
)

type UserController struct {
	userUsecase *usecase.UserUsecase `di.inject:"userUsecase"`
}

func (u *UserController) FindMyAccount(c *fiber.Ctx) error {
	userID := auth.GetUserIDFromClaims(c)

	result, errFind := u.userUsecase.FindUserByID(userID)
	if errFind != nil {
		return errFind
	}

	return c.JSON(result)
}

func (u *UserController) UpdateMyAccount(c *fiber.Ctx) error {
	userID := auth.GetUserIDFromClaims(c)

	payload := new(dto.UpdateUserRequest)
	if err := c.BodyParser(payload); err != nil {
		return exception.NewHTTPError(422, "Failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	result, errUpdate := u.userUsecase.UpdateUser(userID, payload)
	if errUpdate != nil {
		return errUpdate
	}

	return c.JSON(result)
}

func (u *UserController) UpdateMyEmail(c *fiber.Ctx) error {
	userID := auth.GetUserIDFromClaims(c)

	payload := new(dto.UpdateEmailRequest)
	if err := c.BodyParser(payload); err != nil {
		return exception.NewHTTPError(422, "Failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	result, err := u.userUsecase.UpdateEmail(userID, payload)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (u *UserController) UpdateMyPassword(c *fiber.Ctx) error {
	userID := auth.GetUserIDFromClaims(c)

	payload := new(dto.UpdatePasswordRequest)
	if err := c.BodyParser(payload); err != nil {
		return exception.NewHTTPError(422, "Failed to parse body")
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
	userID := auth.GetUserIDFromClaims(c)

	result, err := u.userUsecase.DeleteUser(userID)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
