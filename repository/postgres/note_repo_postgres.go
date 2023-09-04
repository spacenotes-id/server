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
	userID int,
) ([]*sqlc.Note, error) {
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
) ([]*sqlc.Note, error) {
	notes, err := n.querier.FindAllNotesBySpaceID(ctx, int32(spaceID))
	if err != nil {
		log.Error(err)
		return nil, fiber.
			NewError(fiber.StatusInternalServerError, "Failed to find all notes")
	}

	return notes, nil
}

func (n *NoteRepoPostgres) FindAllNotesByStatus(
	ctx context.Context,
	userID int,
	status sqlc.Status,
) ([]*sqlc.Note, error) {
	notes, err := n.querier.FindAllNotesByStatus(ctx, sqlc.FindAllNotesByStatusParams{
		UserID: int32(userID),
		Status: status,
	})
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
	spaceID int,
	status sqlc.Status,
) ([]*sqlc.Note, error) {
	notes, err := n.querier.FindAllNotesBySpaceIDAndStatus(ctx, sqlc.FindAllNotesBySpaceIDAndStatusParams{
		SpaceID: int32(spaceID),
		Status:  status,
	})
	if err != nil {
		log.Error(err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to find all notes by space id and status")
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
