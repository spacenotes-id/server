package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spacenotes-id/server/database/postgres/sqlc"
	"github.com/spacenotes-id/server/dto"
	"github.com/spacenotes-id/server/helper/sql"
	"github.com/spacenotes-id/server/repository/postgres"
)

type NoteUsecase struct {
	noteRepo     *postgres.NoteRepoPostgres `di.inject:"noteRepo"`
	spaceUsecase *SpaceUsecase              `di.inject:"spaceUsecase"`
	userRepo     *postgres.UserRepoPostgres `di.inject:"userRepo"`
}

func (n *NoteUsecase) CreateNote(
	userID int,
	payload *dto.CreateNoteRequest,
) (*dto.CreateNoteResponse, error) {
	ctx := context.Background()

	if _, err := n.userRepo.FindUserByID(ctx, userID); err != nil {
		return nil, err
	}

	if err := n.spaceUsecase.VerifySpaceOwnership(
		userID,
		payload.SpaceID,
	); err != nil {
		return nil, err
	}

	createdNote, errCreate := n.noteRepo.CreateNote(ctx, sqlc.CreateNoteParams{
		UserID:  int32(userID),
		SpaceID: int32(payload.SpaceID),
		Title:   payload.Title,
		Body:    pgtype.Text(sql.NewNullString(payload.Body)),
	})
	if errCreate != nil {
		return nil, errCreate
	}

	response := &dto.CreateNoteResponse{
		Message: "Your new note has been created successfully",
		Data:    *createdNote,
	}

	return response, nil
}

func (n *NoteUsecase) FindAllNotes(
	userID int,
	query map[string]string,
) (*dto.FindAllNotesResponse, error) {
	ctx := context.Background()

	if _, err := n.userRepo.FindUserByID(ctx, userID); err != nil {
		return nil, err
	}

	var notes []*sqlc.Note
	var err error

	if query["status"] != "" {
		notes, err = n.noteRepo.FindAllNotesByStatus(
			ctx,
			userID,
			sqlc.Status(query["status"]),
			query["search"],
		)
		if err != nil {
			return nil, err
		}
	} else {
		notes, err = n.noteRepo.FindAllNotes(ctx, userID, query["search"])
		if err != nil {
			return nil, err
		}
	}

	response := &dto.FindAllNotesResponse{
		Data: notes,
	}

	return response, nil
}

func (n *NoteUsecase) FindAllNotesBySpaceID(
	spaceID int,
	query map[string]string,
) (*dto.FindAllNotesResponse, error) {
	ctx := context.Background()

	if _, err := n.spaceUsecase.FindSpaceByID(spaceID); err != nil {
		return nil, err
	}

	var notes []*sqlc.Note
	var err error

	if query["status"] != "" {
		notes, err = n.noteRepo.FindAllNotesBySpaceIDAndStatus(
			ctx,
			spaceID,
			sqlc.Status(query["status"]),
			query["search"],
		)
		if err != nil {
			return nil, err
		}
	} else {
		notes, err = n.noteRepo.FindAllNotesBySpaceID(ctx, spaceID, query["search"])
		if err != nil {
			return nil, err
		}
	}

	response := &dto.FindAllNotesResponse{
		Data: notes,
	}

	return response, nil
}

func (n *NoteUsecase) VerifyNoteOwnership(userID int, noteID int) error {
	ctx := context.Background()

	note, err := n.noteRepo.FindNoteByID(ctx, noteID)
	if err != nil {
		return err
	}

	if note.UserID != int32(userID) {
		return fiber.
			NewError(fiber.StatusForbidden, "You're not allowed to access this note")
	}

	return nil
}

func (n *NoteUsecase) FindNoteByID(noteID int) (*dto.FindNoteByIDResponse, error) {
	ctx := context.Background()

	note, err := n.noteRepo.FindNoteByID(ctx, noteID)
	if err != nil {
		return nil, err
	}

	response := &dto.FindNoteByIDResponse{
		Data: *note,
	}

	return response, nil
}

func (n *NoteUsecase) UpdateNote(
	userID int,
	noteID int,
	payload *dto.UpdateNoteRequest,
) (*dto.UpdateNoteResponse, error) {
	ctx := context.Background()

	if _, err := n.userRepo.FindUserByID(ctx, userID); err != nil {
		return nil, err
	}

	if _, err := n.noteRepo.FindNoteByID(ctx, noteID); err != nil {
		return nil, err
	}

	if payload.SpaceID != 0 {
		if err := n.spaceUsecase.VerifySpaceOwnership(
			userID,
			payload.SpaceID,
		); err != nil {
			return nil, err
		}
	}

	updatedNote, err := n.noteRepo.UpdateNote(ctx, sqlc.UpdateNoteParams{
		ID:        int32(noteID),
		Title:     pgtype.Text(sql.NewNullString(payload.Title)),
		Body:      pgtype.Text(sql.NewNullString(payload.Body)),
		SpaceID:   pgtype.Int4(sql.NewNullInt(payload.SpaceID)),
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateNoteResponse{
		Message: fmt.Sprintf("Note with id %v has been updated", noteID),
		Data:    *updatedNote,
	}

	return response, nil
}

func (n *NoteUsecase) DeleteNote(noteID int) (*dto.DeleteNoteResponse, error) {
	ctx := context.Background()

	if _, err := n.noteRepo.FindNoteByID(ctx, noteID); err != nil {
		return nil, err
	}

	if err := n.noteRepo.DeleteNote(ctx, noteID); err != nil {
		return nil, err
	}

	response := &dto.DeleteNoteResponse{
		Message: fmt.Sprintf("Note with id %v has been deleted", noteID),
	}

	return response, nil
}

func (n *NoteUsecase) ChangeStatus(
	noteID int,
	status string,
) (*dto.UpdateStatusResponse, error) {
	ctx := context.Background()

	if _, err := n.noteRepo.FindNoteByID(ctx, noteID); err != nil {
		return nil, err
	}

	updatedNote, err := n.noteRepo.UpdateNote(ctx, sqlc.UpdateNoteParams{
		ID:        int32(noteID),
		Status:    sqlc.NullStatus{Status: sqlc.Status(status), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateStatusResponse{
		Message: fmt.Sprintf("Status of note with id %v has been updated", noteID),
		Data:    *updatedNote,
	}

	return response, nil
}
