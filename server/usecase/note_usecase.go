package usecase

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tfkhdyt/SpaceNotes/server/database/postgres/sqlc"
	"github.com/tfkhdyt/SpaceNotes/server/dto"
	"github.com/tfkhdyt/SpaceNotes/server/helper/sql"
	"github.com/tfkhdyt/SpaceNotes/server/repository/postgres"
)

type NoteUsecase struct {
	noteRepo     *postgres.NoteRepoPostgres `di.inject:"noteRepo"`
	spaceUsecase *SpaceUsecase              `di.inject:"spaceUsecase"`
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
