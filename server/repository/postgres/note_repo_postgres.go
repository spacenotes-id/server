package postgres

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/SpaceNotes/server/database/postgres/sqlc"
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
