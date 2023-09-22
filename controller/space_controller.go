package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"

	"github.com/spacenotes-id/server/dto"
	"github.com/spacenotes-id/server/helper/auth"
	"github.com/spacenotes-id/server/helper/exception"
	"github.com/spacenotes-id/server/helper/space"
	"github.com/spacenotes-id/server/usecase"
)

type SpaceController struct {
	spaceUsecase *usecase.SpaceUsecase `di.inject:"spaceUsecase"`
}

func (s *SpaceController) CreateSpace(c *fiber.Ctx) error {
	userID, err := auth.GetUserIDFromClaims(c)
	if err != nil {
		return err
	}

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
	userID, err := auth.GetUserIDFromClaims(c)
	if err != nil {
		return err
	}

	response, err := s.spaceUsecase.FindAllSpacesByUserID(userID)
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func (s *SpaceController) FindSpaceByID(c *fiber.Ctx) error {
	spaceID, err := space.GetSpaceIDFromParams(c)
	if err != nil {
		return err
	}

	response, errFind := s.spaceUsecase.FindSpaceByID(spaceID)
	if errFind != nil {
		return errFind
	}

	return c.JSON(response)
}

func (s *SpaceController) UpdateSpace(c *fiber.Ctx) error {
	spaceID, err := space.GetSpaceIDFromParams(c)
	if err != nil {
		return err
	}

	payload := new(dto.UpdateSpaceRequest)
	if err := c.BodyParser(payload); err != nil {
		return fiber.
			NewError(fiber.StatusUnprocessableEntity, "Failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	response, errUpdate := s.spaceUsecase.UpdateSpace(spaceID, payload)
	if errUpdate != nil {
		return errUpdate
	}

	return c.JSON(response)
}

func (s *SpaceController) DeleteSpace(c *fiber.Ctx) error {
	spaceID, err := space.GetSpaceIDFromParams(c)
	if err != nil {
		return err
	}

	response, errDelete := s.spaceUsecase.DeleteSpace(spaceID)
	if errDelete != nil {
		return errDelete
	}

	return c.JSON(response)
}
