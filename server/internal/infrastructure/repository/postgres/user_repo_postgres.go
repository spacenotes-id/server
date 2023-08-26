package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tfkhdyt/SpaceNotes/server/internal/domain/entity"
	"github.com/tfkhdyt/SpaceNotes/server/internal/infrastructure/database/postgres/sqlc"
	"github.com/tfkhdyt/SpaceNotes/server/pkg/exception"
)

type UserRepoPostgres struct {
	querier sqlc.Querier `di.inject:"querier"`
}

func (u *UserRepoPostgres) CreateUser(
	ctx context.Context,
	user *entity.NewUser,
) (*entity.CreatedUser, error) {
	result, err := u.querier.CreateUser(ctx, sqlc.CreateUserParams{
		FullName: pgtype.Text(user.FullName),
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		log.Printf("ERROR(CreateUser): %v\n", err)
		return nil, exception.NewHTTPError(500, "Failed to create new user")
	}

	return &entity.CreatedUser{
		ID:        int(result.ID),
		FullName:  sql.NullString(result.FullName),
		Username:  result.Username,
		Email:     result.Email,
		CreatedAt: result.CreatedAt.Time,
	}, nil
}

func (u *UserRepoPostgres) FindUserByID(
	ctx context.Context,
	id int,
) (*entity.User, error) {
	user, err := u.querier.FindUserByID(ctx, int32(id))
	if err != nil {
		return nil, exception.NewHTTPError(404, fmt.Sprintf(
			"User with id %v is not found",
			id,
		))
	}

	return &entity.User{
		ID:        int(user.ID),
		FullName:  sql.NullString(user.FullName),
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}, nil
}

func (u *UserRepoPostgres) FindUserByUsername(
	ctx context.Context,
	username string,
) (*entity.User, error) {
	user, err := u.querier.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, exception.NewHTTPError(404, fmt.Sprintf(
			"User with username %v is not found",
			username,
		))
	}

	return &entity.User{
		ID:        int(user.ID),
		FullName:  sql.NullString(user.FullName),
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}, nil
}

func (u *UserRepoPostgres) FindUserByEmail(
	ctx context.Context,
	email string,
) (*entity.User, error) {
	user, err := u.querier.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, exception.NewHTTPError(404, fmt.Sprintf(
			"User with email %v is not found",
			email,
		))
	}

	return &entity.User{
		ID:        int(user.ID),
		FullName:  sql.NullString(user.FullName),
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}, nil
}

func (u *UserRepoPostgres) UpdateUser(
	ctx context.Context,
	id int,
	data *entity.UpdateUser,
) (*entity.UpdatedUser, error) {
	updatedUser, err := u.querier.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:       int32(id),
		FullName: pgtype.Text(data.FullName),
		Username: pgtype.Text(data.Username),
	})
	if err != nil {
		log.Printf("ERROR(UpdateUser): %v\n", err)
		return nil, exception.NewHTTPError(500, fmt.Sprintf(
			"Failed to update user with id %v",
			id,
		))
	}

	return &entity.UpdatedUser{
		ID:        int(updatedUser.ID),
		FullName:  sql.NullString(updatedUser.FullName),
		Username:  updatedUser.Username,
		Email:     updatedUser.Email,
		CreatedAt: updatedUser.CreatedAt.Time,
		UpdatedAt: updatedUser.UpdatedAt.Time,
	}, nil
}

func (u *UserRepoPostgres) UpdateEmail(
	ctx context.Context,
	id int,
	email string,
) (*entity.UpdatedUser, error) {
	updatedUser, err := u.querier.UpdateEmail(ctx, sqlc.UpdateEmailParams{
		ID:    int32(id),
		Email: email,
	})
	if err != nil {
		log.Printf("ERROR(UpdateUser): %v\n", err)
		return nil, exception.NewHTTPError(500, fmt.Sprintf(
			"Failed to update email with user id %v",
			id,
		))
	}

	return &entity.UpdatedUser{
		ID:        int(updatedUser.ID),
		FullName:  sql.NullString(updatedUser.FullName),
		Username:  updatedUser.Username,
		Email:     updatedUser.Email,
		CreatedAt: updatedUser.CreatedAt.Time,
		UpdatedAt: updatedUser.UpdatedAt.Time,
	}, nil
}

func (u *UserRepoPostgres) UpdatePassword(
	ctx context.Context,
	id int,
	password string,
) error {
	if err := u.querier.UpdatePassword(ctx, sqlc.UpdatePasswordParams{
		ID:       int32(id),
		Password: password,
	}); err != nil {
		log.Printf("ERROR(UpdateUser): %v\n", err)
		return exception.NewHTTPError(500,
			fmt.Sprintf(
				"Failed to update password with user id %v",
				id,
			))
	}

	return nil
}

func (u *UserRepoPostgres) DeleteUser(ctx context.Context, id int) error {
	if err := u.querier.DeleteUser(ctx, int32(id)); err != nil {
		log.Printf("ERROR(UpdateUser): %v\n", err)
		return exception.NewHTTPError(500,
			fmt.Sprintf(
				"Failed to delete user with id %v",
				id,
			))
	}

	return nil
}
