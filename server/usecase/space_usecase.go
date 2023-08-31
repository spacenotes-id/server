package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tfkhdyt/SpaceNotes/server/database/postgres/sqlc"
	"github.com/tfkhdyt/SpaceNotes/server/dto"
	"github.com/tfkhdyt/SpaceNotes/server/helper/sql"
	"github.com/tfkhdyt/SpaceNotes/server/repository/postgres"
)

type SpaceUsecase struct {
	spaceRepo *postgres.SpaceRepoPostgres `di.inject:"spaceRepo"`
	userRepo  *postgres.UserRepoPostgres  `di.inject:"userRepo"`
}

func (s *SpaceUsecase) CreateSpace(userID int, payload *dto.CreateSpaceRequest) (*dto.CreateSpaceResponse, error) {
	ctx := context.Background()

	if _, err := s.userRepo.FindUserByID(ctx, userID); err != nil {
		return nil, err
	}

	createdSpace, err := s.spaceRepo.CreateSpace(ctx, sqlc.CreateSpaceParams{
		Name:     payload.Name,
		Emoji:    pgtype.Text(sql.NewNullString(payload.Emoji)),
		IsLocked: payload.IsLocked,
		UserID:   int32(userID),
	})
	if err != nil {
		return nil, err
	}

	response := &dto.CreateSpaceResponse{
		Message: "Your new space has been created successfully",
		Data:    *createdSpace,
	}

	return response, nil
}

func (s *SpaceUsecase) FindAllSpacesByUserID(
	userID int,
) (*dto.FindAllSpacesByUserIDResponse, error) {
	ctx := context.Background()

	if _, err := s.userRepo.FindUserByID(ctx, userID); err != nil {
		return nil, err
	}

	spaces, err := s.spaceRepo.FindAllSpacesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	response := &dto.FindAllSpacesByUserIDResponse{
		Data: spaces,
	}

	return response, nil
}

func (s *SpaceUsecase) FindSpaceByID(
	spaceID int,
) (*dto.FindSpaceByIDResponse, error) {
	ctx := context.Background()

	space, err := s.spaceRepo.FindSpaceByID(ctx, spaceID)
	if err != nil {
		return nil, err
	}

	response := &dto.FindSpaceByIDResponse{
		Data: *space,
	}

	return response, nil
}

func (s *SpaceUsecase) VerifySpaceOwnership(userID int, spaceID int) error {
	ctx := context.Background()

	space, err := s.spaceRepo.FindSpaceByID(ctx, spaceID)
	if err != nil {
		return err
	}

	if space.UserID != int32(userID) {
		return fiber.NewError(
			fiber.StatusForbidden,
			"You're not allowed to access this space",
		)
	}

	return nil
}

func (s *SpaceUsecase) UpdateSpace(
	id int,
	data *dto.UpdateSpaceRequest,
) (*dto.UpdateSpaceResponse, error) {
	ctx := context.Background()

	if _, err := s.spaceRepo.FindSpaceByID(ctx, id); err != nil {
		return nil, err
	}

	updatedUser, err := s.spaceRepo.UpdateSpace(ctx, sqlc.UpdateSpaceParams{
		ID:        int32(id),
		Name:      pgtype.Text(sql.NewNullString(data.Name)),
		Emoji:     pgtype.Text(sql.NewNullString(data.Emoji)),
		IsLocked:  pgtype.Bool(sql.NewNullBool(&data.IsLocked)),
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateSpaceResponse{
		Message: fmt.Sprintf(
			"Space with id %v has been updated successfully",
			updatedUser.ID,
		),
		Data: *updatedUser,
	}

	return response, nil
}

func (s *SpaceUsecase) DeleteSpace(id int) (*dto.DeleteSpaceResponse, error) {
	ctx := context.Background()

	space, err := s.spaceRepo.FindSpaceByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if space.IsLocked {
		return nil, fiber.NewError(
			fiber.StatusBadRequest,
			"This space cannot be deleted because it's locked",
		)
	}

	if err := s.spaceRepo.DeleteSpace(ctx, id); err != nil {
		return nil, err
	}

	response := &dto.DeleteSpaceResponse{
		Message: fmt.Sprintf("Space with id %v has been deleted successfully", id),
	}

	return response, nil
}
