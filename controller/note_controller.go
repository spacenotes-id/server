package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spacenotes-id/server/dto"
	"github.com/spacenotes-id/server/helper/auth"
	"github.com/spacenotes-id/server/helper/note"
	"github.com/spacenotes-id/server/helper/space"
	"github.com/spacenotes-id/server/helper/validation"
	"github.com/spacenotes-id/server/usecase"
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

func (n *NoteController) FindAllTrashedNotes(c *fiber.Ctx) error {
	userID := auth.GetUserIDFromClaims(c)

	notes, err := n.noteUsecase.FindAllTrashedNotes(userID)
	if err != nil {
		return err
	}

	return c.JSON(notes)
}

func (n *NoteController) FindAllFavoriteNotes(c *fiber.Ctx) error {
	userID := auth.GetUserIDFromClaims(c)

	notes, err := n.noteUsecase.FindAllFavoriteNotes(userID)
	if err != nil {
		return err
	}

	return c.JSON(notes)
}

func (n *NoteController) FindAllArchivedNotes(c *fiber.Ctx) error {
	userID := auth.GetUserIDFromClaims(c)

	notes, err := n.noteUsecase.FindAllArchivedNotes(userID)
	if err != nil {
		return err
	}

	return c.JSON(notes)
}

func (n *NoteController) FindNoteByID(c *fiber.Ctx) error {
	noteID, err := note.GetNoteIDFromParams(c)
	if err != nil {
		return err
	}

	note, errFind := n.noteUsecase.FindNoteByID(noteID)
	if errFind != nil {
		return errFind
	}

	return c.JSON(note)
}

func (n *NoteController) UpdateNote(c *fiber.Ctx) error {
	userID := auth.GetUserIDFromClaims(c)
	noteID, err := note.GetNoteIDFromParams(c)
	if err != nil {
		return err
	}

	payload := new(dto.UpdateNoteRequest)
	if err := validation.ValidateBody(c, payload); err != nil {
		return err
	}

	updatedNote, errUpdate := n.noteUsecase.UpdateNote(userID, noteID, payload)
	if errUpdate != nil {
		return errUpdate
	}

	return c.JSON(updatedNote)
}

func (n *NoteController) DeleteNote(c *fiber.Ctx) error {
	noteID, err := note.GetNoteIDFromParams(c)
	if err != nil {
		return err
	}

	response, errDelete := n.noteUsecase.DeleteNote(noteID)
	if errDelete != nil {
		return errDelete
	}

	return c.JSON(response)
}
