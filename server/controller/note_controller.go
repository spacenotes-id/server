package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/SpaceNotes/server/dto"
	"github.com/tfkhdyt/SpaceNotes/server/helper/auth"
	"github.com/tfkhdyt/SpaceNotes/server/helper/space"
	"github.com/tfkhdyt/SpaceNotes/server/helper/validation"
	"github.com/tfkhdyt/SpaceNotes/server/usecase"
)

type NoteController struct {
	noteUsecase *usecase.NoteUsecase `di.inject:"noteUsecase"`
}

func (n *NoteController) CreateNote(c *fiber.Ctx) error {
	userID := auth.GetUserIDFromClaims(c)

	payload := new(dto.CreateNoteRequest)
	if err := validation.ValidateBody(c, payload); err != nil {
		return err
	}

	createdNote, err := n.noteUsecase.CreateNote(userID, payload)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(createdNote)
}

func (n *NoteController) FindAllNotes(c *fiber.Ctx) error {
	userID := auth.GetUserIDFromClaims(c)

	notes, err := n.noteUsecase.FindAllNotes(userID)
	if err != nil {
		return err
	}

	return c.JSON(notes)
}

func (n *NoteController) FindAllNotesBySpaceID(c *fiber.Ctx) error {
	spaceID, err := space.GetSpaceIDFromParams(c)
	if err != nil {
		return err
	}

	notes, errFind := n.noteUsecase.FindAllNotesBySpaceID(spaceID)
	if errFind != nil {
		return errFind
	}

	return c.JSON(notes)
}
