package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/SpaceNotes/server/dto"
	"github.com/tfkhdyt/SpaceNotes/server/helper/auth"
	"github.com/tfkhdyt/SpaceNotes/server/helper/exception"
	"github.com/tfkhdyt/SpaceNotes/server/usecase"
)

type SpaceController struct {
	spaceUsecase *usecase.SpaceUsecase `di.inject:"spaceUsecase"`
}

func (s *SpaceController) CreateSpace(c *fiber.Ctx) error {
	userID := auth.GetUserIDFromClaims(c)

	payload := new(dto.CreateSpaceRequest)
	if err := c.BodyParser(payload); err != nil {
		return fiber.
			NewError(fiber.StatusUnprocessableEntity, "Failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	response, err := s.spaceUsecase.CreateSpace(userID, payload)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (s *SpaceController) FindAllSpacesByUserID(c *fiber.Ctx) error {
	userID := auth.GetUserIDFromClaims(c)

	response, err := s.spaceUsecase.FindAllSpacesByUserID(userID)
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func (s *SpaceController) FindSpaceByID(c *fiber.Ctx) error {
	spaceID, err := c.ParamsInt("space_id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to get space id")
	}

	response, err := s.spaceUsecase.FindSpaceByID(spaceID)
	if err != nil {
		return err
	}

	return c.JSON(response)
}