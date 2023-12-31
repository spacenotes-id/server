package postgres

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spacenotes-id/server/database/postgres/sqlc"
)

type SpaceRepoPostgres struct {
	querier *sqlc.Queries `di.inject:"querier"`
}

func (s *SpaceRepoPostgres) CreateSpace(
	ctx context.Context,
	data sqlc.CreateSpaceParams,
) (*sqlc.CreateSpaceRow, error) {
	createdSpace, err := s.querier.CreateSpace(ctx, data)
	if err != nil {
		log.Error(err)
		return nil, fiber.
			NewError(fiber.StatusInternalServerError, "Failed to create new space")
	}

	return createdSpace, nil
}

func (s *SpaceRepoPostgres) FindAllSpacesByUserID(
	ctx context.Context,
	userID int,
) ([]*sqlc.FindAllSpacesByUserIDRow, error) {
	spaces, err := s.querier.FindAllSpacesByUserID(ctx, int32(userID))
	if err != nil {
		log.Error(err)
		return nil, fiber.
			NewError(fiber.StatusInternalServerError, "Failed to get all spaces")
	}

	return spaces, nil
}

func (s *SpaceRepoPostgres) FindSpaceByID(
	ctx context.Context,
	spaceID int,
) (*sqlc.Space, error) {
	space, err := s.querier.FindSpaceByID(ctx, int32(spaceID))
	if err != nil {
		return nil, fiber.NewError(
			fiber.StatusNotFound,
			fmt.Sprintf("Space with id %v is not found", spaceID),
		)
	}

	return space, nil
}

func (s *SpaceRepoPostgres) UpdateSpace(
	ctx context.Context,
	data sqlc.UpdateSpaceParams,
) (*sqlc.UpdateSpaceRow, error) {
	updatedSpace, err := s.querier.UpdateSpace(ctx, data)
	if err != nil {
		log.Error(err)
		return nil, fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("Failed to update space with id %v", data.ID),
		)
	}

	return updatedSpace, nil
}

func (s *SpaceRepoPostgres) DeleteSpace(ctx context.Context, id int) error {
	if err := s.querier.DeleteSpace(ctx, int32(id)); err != nil {
		log.Error(err)
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("Failed to delete space with id %v", id),
		)
	}

	return nil
}
