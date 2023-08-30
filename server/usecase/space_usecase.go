package usecase

import (
	"context"

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

func (s *SpaceUsecase) verifyNameAvailability(
	ctx context.Context,
	name string,
) error {
	if _, err := s.spaceRepo.FindSpaceByName(ctx, name); err == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Name has been used")
	}

	return nil
}

func (s *SpaceUsecase) CreateSpace(userID int, payload *dto.CreateSpaceRequest) (*dto.CreateSpaceResponse, error) {
	ctx := context.Background()

	if _, err := s.userRepo.FindUserByID(ctx, userID); err != nil {
		return nil, err
	}

	if err := s.verifyNameAvailability(ctx, payload.Name); err != nil {
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

func (s SpaceUsecase) VerifySpaceOwnership(userID int, spaceID int) error {
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
