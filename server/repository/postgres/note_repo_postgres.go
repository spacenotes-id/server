package postgres

import (
	"context"

	"github.com/gofiber/fiber/v2"
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
		return nil, fiber.
			NewError(fiber.StatusInternalServerError, "Failed to find all trashed notes")
	}

	return notes, nil
}
