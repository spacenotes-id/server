package postgres

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spacenotes-id/SpaceNotes/server/database/postgres"
	"github.com/spacenotes-id/SpaceNotes/server/database/postgres/sqlc"
)

type UserRepoPostgres struct {
	querier *sqlc.Queries `di.inject:"querier"`
}

func (u *UserRepoPostgres) CreateUser(
	ctx context.Context,
	newUser sqlc.CreateUserParams,
) (*sqlc.CreateUserRow, error) {
	tx, err := postgres.Pool.Begin(ctx)
	if err != nil {
		log.Error(err)
		return nil, fiber.
			NewError(fiber.StatusInternalServerError, "Failed to start transaction")
	}
	// defer tx.Rollback(ctx)
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	qtx := u.querier.WithTx(tx)

	result, err := qtx.CreateUser(ctx, newUser)
	if err != nil {
		log.Error(err)
		return nil, fiber.
			NewError(fiber.StatusInternalServerError, "Failed to create new user")
	}

	if _, errSpace := qtx.CreateSpace(ctx, sqlc.CreateSpaceParams{
		Name:     "General",
		Emoji:    pgtype.Text{String: "1F3E2", Valid: true},
		IsLocked: true,
		UserID:   result.ID,
	}); errSpace != nil {
		log.Error(err)
		return nil, fiber.NewError(
			fiber.StatusInternalServerError,
			"Failed to create general space",
		)
	}

	if err := tx.Commit(ctx); err != nil {
		log.Error(err)
		return nil, fiber.
			NewError(fiber.StatusInternalServerError, "Failed to commit transaction")
	}

	return result, nil
}

func (u *UserRepoPostgres) FindUserByID(
	ctx context.Context,
	id int,
) (*sqlc.User, error) {
	user, err := u.querier.FindUserByID(ctx, int32(id))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, fmt.Sprintf(
			"User with id %v is not found",
			id,
		))
	}

	return user, nil
}

func (u *UserRepoPostgres) FindUserByUsername(
	ctx context.Context,
	username string,
) (*sqlc.User, error) {
	user, err := u.querier.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, fmt.Sprintf(
			"User with username %v is not found",
			username,
		))
	}

	return user, nil
}

func (u *UserRepoPostgres) FindUserByEmail(
	ctx context.Context,
	email string,
) (*sqlc.User, error) {
	user, err := u.querier.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, fmt.Sprintf(
			"User with email %v is not found",
			email,
		))
	}

	return user, nil
}

func (u *UserRepoPostgres) UpdateUser(
	ctx context.Context,
	data sqlc.UpdateUserParams,
) (*sqlc.UpdateUserRow, error) {
	updatedUser, err := u.querier.UpdateUser(ctx, data)
	if err != nil {
		log.Error(err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf(
			"Failed to update user with id %v",
			data.ID,
		))
	}

	return updatedUser, nil
}

func (u *UserRepoPostgres) UpdateEmail(
	ctx context.Context,
	data sqlc.UpdateEmailParams,
) (*sqlc.UpdateEmailRow, error) {
	updatedUser, err := u.querier.UpdateEmail(ctx, data)
	if err != nil {
		log.Error(err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf(
			"Failed to update email with user id %v",
			data.ID,
		))
	}

	return updatedUser, nil
}

func (u *UserRepoPostgres) UpdatePassword(
	ctx context.Context,
	data sqlc.UpdatePasswordParams,
) error {
	if err := u.querier.UpdatePassword(ctx, data); err != nil {
		log.Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf(
			"Failed to update password with user id %v",
			data.ID,
		))
	}

	return nil
}

func (u *UserRepoPostgres) DeleteUser(ctx context.Context, id int) error {
	if err := u.querier.DeleteUser(ctx, int32(id)); err != nil {
		log.Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf(
			"Failed to delete user with id %v",
			id,
		))
	}

	return nil
}
