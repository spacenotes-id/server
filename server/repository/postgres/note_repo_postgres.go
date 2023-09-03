package postgres

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spacenotes-id/SpaceNotes/server/database/postgres/sqlc"
)

type NoteRepoPostgres struct {
	querier *sqlc.Queries `di.inject:"querier"`
}

func (n *NoteRepoPostgres) CreateNote(
	ctx context.Context,
	data sqlc.CreateNoteParams,
) (*sqlc.CreateNoteRow, error) {
	createdNote, err := n.querier.CreateNote(ctx, data)
	if err != nil {
		log.Error(err)
		return nil, fiber.
			NewError(fiber.StatusInternalServerError, "Failed to create new note")
	}

	return createdNote, nil
}

func (n *NoteRepoPostgres) FindAllNotes(
	ctx context.Context,
	userID int,
) ([]*sqlc.FindAllNotesRow, error) {
	notes, err := n.querier.FindAllNotes(ctx, int32(userID))
	if err != nil {
		log.Error(err)
		return nil, fiber.
			NewError(fiber.StatusInternalServerError, "Failed to find all notes")
	}

	return notes, nil
}

func (n *NoteRepoPostgres) FindAllNotesBySpaceID(
	ctx context.Context,
	spaceID int,
) ([]*sqlc.FindAllNotesBySpaceIDRow, error) {
	notes, err := n.querier.FindAllNotesBySpaceID(ctx, int32(spaceID))
	if err != nil {
		log.Error(err)
		return nil, fiber.
			NewError(fiber.StatusInternalServerError, "Failed to find all notes")
	}

	return notes, nil
}

func (n *NoteRepoPostgres) FindAllTrashedNotes(
	ctx context.Context,
	userID int,
) ([]*sqlc.FindAllTrashedNotesRow, error) {
	notes, err := n.querier.FindAllTrashedNotes(ctx, int32(userID))
	if err != nil {
		log.Error(err)
		return nil, fiber.NewError(
			fiber.StatusInternalServerError,
			"Failed to find all trashed notes",
		)
	}

	return notes, nil
}

func (n *NoteRepoPostgres) FindAllFavoriteNotes(
	ctx context.Context,
	userID int,
) ([]*sqlc.FindAllFavoriteNotesRow, error) {
	notes, err := n.querier.FindAllFavoriteNotes(ctx, int32(userID))
	if err != nil {
		log.Error(err)
		return nil, fiber.NewError(
			fiber.StatusInternalServerError,
			"Failed to find all favorite notes",
		)
	}

	return notes, nil
}

func (n *NoteRepoPostgres) FindAllArchivedNotes(
	ctx context.Context,
	userID int,
) ([]*sqlc.FindAllArchivedNotesRow, error) {
	notes, err := n.querier.FindAllArchivedNotes(ctx, int32(userID))
	if err != nil {
		log.Error(err)
		return nil, fiber.NewError(
			fiber.StatusInternalServerError,
			"Failed to find all archived notes",
		)
	}

	return notes, nil
}

func (n *NoteRepoPostgres) FindNoteByID(
	ctx context.Context,
	id int,
) (*sqlc.Note, error) {
	note, err := n.querier.FindNoteByID(ctx, int32(id))
	if err != nil {
		return nil, fiber.NewError(
			fiber.StatusNotFound,
			fmt.Sprintf("Note with id %v is not found", id),
		)
	}

	return note, nil
}

func (n *NoteRepoPostgres) UpdateNote(
	ctx context.Context,
	data sqlc.UpdateNoteParams,
) (*sqlc.UpdateNoteRow, error) {
	updatedNote, err := n.querier.UpdateNote(ctx, data)
	if err != nil {
		log.Error(err)
		return nil, fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("Failed to to update note with id %v", data.ID),
		)
	}

	return updatedNote, nil
}
