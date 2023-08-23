package repository

import (
	"context"

	"github.com/tfkhdyt/SpaceNotes/server/internal/domain/entity"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *entity.NewUser) (*entity.CreatedUser, error)

	FindUserByID(ctx context.Context, id int) (*entity.User, error)
	FindUserByUsername(ctx context.Context, username string) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)

	UpdateUser(ctx context.Context, id int, data *entity.UpdateUser) (*entity.UpdatedUser, error)
	UpdateEmail(ctx context.Context, id int, email string) (*entity.UpdatedUser, error)
	UpdatePassword(ctx context.Context, id int, password string) error

	DeleteUser(ctx context.Context, id int) error
}
