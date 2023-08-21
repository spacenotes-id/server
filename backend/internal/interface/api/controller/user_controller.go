package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/SpaceNotes/backend/internal/application/dto"
	"github.com/tfkhdyt/SpaceNotes/backend/internal/application/usecase"
	"github.com/tfkhdyt/SpaceNotes/backend/pkg/exception"
)

type UserController struct {
	userUsecase *usecase.UserUsecase `di.inject:"userUsecase"`
}

func (u *UserController) FindUserByID(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("userID")
	if err != nil {
		return exception.NewHTTPError(400, "Invalid user id")
	}

	result, errFind := u.userUsecase.FindUserByID(userID)
	if errFind != nil {
		return errFind
	}

	return c.JSON(result)
}

func (u *UserController) UpdateUser(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("userID")
	if err != nil {
		return exception.NewHTTPError(400, "Invalid user id")
	}

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
