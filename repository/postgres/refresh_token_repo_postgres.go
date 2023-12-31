package postgres

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spacenotes-id/server/database/postgres/sqlc"
)

type RefreshTokenRepoPostgres struct {
	querier *sqlc.Queries `di.inject:"querier"`
}

func (r *RefreshTokenRepoPostgres) AddToken(
	ctx context.Context,
	token string,
) error {
	if err := r.querier.AddToken(ctx, token); err != nil {
		log.Error(err)
		return fiber.
			NewError(fiber.StatusInternalServerError, "Failed to add refresh token")
	}

	return nil
}

func (r *RefreshTokenRepoPostgres) VerifyToken(
	ctx context.Context,
	token string,
) error {
	if _, err := r.querier.FindToken(ctx, token); err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Token is not found")
	}

	return nil
}

func (r *RefreshTokenRepoPostgres) DeleteToken(
	ctx context.Context,
	token string,
) error {
	if err := r.querier.DeleteToken(ctx, token); err != nil {
		log.Error(err)
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Failed to delete refresh token",
		)
	}

	return nil
}
