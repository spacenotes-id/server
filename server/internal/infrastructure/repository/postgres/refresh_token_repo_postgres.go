package postgres

import (
	"context"

	"github.com/tfkhdyt/SpaceNotes/server/internal/domain/entity"
	"github.com/tfkhdyt/SpaceNotes/server/internal/infrastructure/database/postgres/sqlc"
	"github.com/tfkhdyt/SpaceNotes/server/pkg/exception"
)

type RefreshTokenRepoPostgres struct {
	querier sqlc.Querier `di.inject:"querier"`
}

func (r *RefreshTokenRepoPostgres) AddToken(ctx context.Context, token string) (*entity.RefreshToken, error) {
	token, err := r.querier.AddToken(ctx, token)
	if err != nil {
		return nil, exception.NewHTTPError(500, "Failed to add refresh token")
	}

	return &entity.RefreshToken{
		Token: token,
	}, nil
}

func (r *RefreshTokenRepoPostgres) FindToken(ctx context.Context, token string) (*entity.RefreshToken, error) {
	token, err := r.querier.FindToken(ctx, token)
	if err != nil {
		return nil, exception.NewHTTPError(404, "Token is not found")
	}

	return &entity.RefreshToken{
		Token: token,
	}, nil
}

func (r *RefreshTokenRepoPostgres) DeleteToken(ctx context.Context, token string) error {
	if err := r.querier.DeleteToken(ctx, token); err != nil {
		return exception.NewHTTPError(500, "Failed to delete refresh token")
	}

	return nil
}
