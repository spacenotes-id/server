package postgres

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/tfkhdyt/SpaceNotes/server/database/postgres/sqlc"
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

func (s *SpaceRepoPostgres) FindSpaceByName(
	ctx context.Context,
	name string,
) (*sqlc.Space, error) {
	space, err := s.querier.FindSpaceByName(ctx, name)
	if err != nil {
		return nil, fiber.NewError(
			fiber.StatusNotFound,
			fmt.Sprintf("Space with name %s is not found", name),
		)
	}

	return space, nil
}

func (s *SpaceRepoPostgres) FindAllSpacesByUserID(
	ctx context.Context,
	userID int,
) ([]*sqlc.FindAllSpacesByUserIDRow, error) {
	spaces, err := s.querier.FindAllSpacesByUserID(ctx, int32(userID))
	if err != nil {
		return nil, fiber.
			NewError(fiber.StatusInternalServerError, "Failed to get all spaces")
	}

	return spaces, nil
}
