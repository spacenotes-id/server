package postgres

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spacenotes-id/server/database/postgres/sqlc"
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
	data sqlc.FindAllNotesParams,
) ([]*sqlc.Note, error) {
	notes, err := n.querier.FindAllNotes(ctx, data)
	if err != nil {
		log.Error(err)
		return nil, fiber.
			NewError(fiber.StatusInternalServerError, "Failed to find all notes")
	}

	return notes, nil
}

func (n *NoteRepoPostgres) FindAllNotesBySpaceID(
	ctx context.Context,
	data sqlc.FindAllNotesBySpaceIDParams,
) ([]*sqlc.Note, error) {
	notes, err := n.querier.FindAllNotesBySpaceID(ctx, data)
	if err != nil {
		log.Error(err)
		return nil, fiber.
			NewError(fiber.StatusInternalServerError, "Failed to find all notes")
	}

	return notes, nil
}

func (n *NoteRepoPostgres) FindAllNotesByStatus(
	ctx context.Context,
	data sqlc.FindAllNotesByStatusParams,
) ([]*sqlc.Note, error) {
	notes, err := n.querier.FindAllNotesByStatus(ctx, data)
	if err != nil {
		log.Error(err)
		return nil, fiber.NewError(
			fiber.StatusInternalServerError,
			"Failed to find all notes by status",
		)
	}

	return notes, nil
}

func (n *NoteRepoPostgres) FindAllNotesBySpaceIDAndStatus(
	ctx context.Context,
	data sqlc.FindAllNotesBySpaceIDAndStatusParams,
) ([]*sqlc.Note, error) {
	notes, err := n.querier.FindAllNotesBySpaceIDAndStatus(ctx, data)
	if err != nil {
		log.Error(err)
		return nil, fiber.NewError(
			fiber.StatusInternalServerError,
			"Failed to find all notes by space id and status",
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

func (n *NoteRepoPostgres) DeleteNote(ctx context.Context, id int) error {
	if err := n.querier.DeleteNote(ctx, int32(id)); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("Failed to delete note with id %v", id),
		)
	}

	return nil
}
