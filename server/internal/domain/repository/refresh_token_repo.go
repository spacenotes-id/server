package repository

import (
	"context"

	"github.com/tfkhdyt/SpaceNotes/server/internal/domain/entity"
)

type RefreshTokenRepo interface {
	AddToken(ctx context.Context, token string) (*entity.RefreshToken, error)
	FindToken(ctx context.Context, token string) (*entity.RefreshToken, error)
	DeleteToken(ctx context.Context, token string) error
}
