package usecase

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spacenotes-id/SpaceNotes/server/database/postgres/sqlc"
	"github.com/spacenotes-id/SpaceNotes/server/dto"
	"github.com/spacenotes-id/SpaceNotes/server/helper/sql"
	"github.com/spacenotes-id/SpaceNotes/server/repository/postgres"
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
) (*dto.FindAllNotesResponse, error) {
	ctx := context.Background()

	if _, err := n.userRepo.FindUserByID(ctx, userID); err != nil {
		return nil, err
	}

	notes, err := n.noteRepo.FindAllNotes(ctx, userID)
	if err != nil {
		return nil, err
	}

	response := &dto.FindAllNotesResponse{
		Data: notes,
	}

	return response, nil
}

func (n *NoteUsecase) FindAllNotesBySpaceID(
	spaceID int,
) (*dto.FindAllNotesBySpaceIDResponse, error) {
	ctx := context.Background()

	if _, err := n.spaceUsecase.FindSpaceByID(spaceID); err != nil {
		return nil, err
	}

	notes, err := n.noteRepo.FindAllNotesBySpaceID(ctx, spaceID)
	if err != nil {
		return nil, err
	}

	response := &dto.FindAllNotesBySpaceIDResponse{
		Data: notes,
	}

	return response, nil
}
