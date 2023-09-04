package usecase

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spacenotes-id/server/database/postgres/sqlc"
	"github.com/spacenotes-id/server/dto"
	"github.com/spacenotes-id/server/helper/sql"
	"github.com/spacenotes-id/server/repository/postgres"
	"github.com/spacenotes-id/server/service"
)

type UserUsecase struct {
	userRepo      *postgres.UserRepoPostgres `di.inject:"userRepo"`
	bcryptService *service.BcryptService     `di.inject:"bcryptService"`
}

func (u *UserUsecase) FindUserByID(
	id int,
) (*dto.FindUserByIDResponse, error) {
	ctx := context.Background()

	user, err := u.userRepo.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	response := &dto.FindUserByIDResponse{
		Data: *user,
	}

	return response, nil
}

func (a *UserUsecase) verifyEmailAvailability(
	ctx context.Context,
	email string,
) error {
	if _, err := a.userRepo.FindUserByEmail(ctx, email); err == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Email has been used")
	}

	return nil
}

func (u *UserUsecase) verifyUsernameAvailability(
	ctx context.Context,
	username string,
) error {
	if _, err := u.userRepo.FindUserByUsername(
		ctx,
		username,
	); err == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Username has been used")
	}

	return nil
}

func (u *UserUsecase) UpdateUser(
	id int,
	data *dto.UpdateUserRequest,
) (*dto.UpdateUserResponse, error) {
	ctx := context.Background()

	if _, err := u.userRepo.FindUserByID(ctx, id); err != nil {
		return nil, err
	}

	if data.Username != "" {
		if err := u.verifyUsernameAvailability(ctx, data.Username); err != nil {
			return nil, err
		}
	}

	updatedUser, err := u.userRepo.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:        int32(id),
		FullName:  pgtype.Text(sql.NewNullString(data.FullName)),
		Username:  pgtype.Text(sql.NewNullString(data.Username)),
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateUserResponse{
		Message: "Your account data has been updated successfully ",
		Data:    *updatedUser,
	}

	return response, nil
}

func (u *UserUsecase) UpdateEmail(
	id int,
	data *dto.UpdateEmailRequest,
) (*dto.UpdateUserResponse, error) {
	ctx := context.Background()

	user, err := u.userRepo.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := u.verifyEmailAvailability(
		ctx,
		data.NewEmail,
	); err != nil {
		return nil, err
	}

	if err := u.bcryptService.ComparePassword(
		user.Password,
		data.Password,
	); err != nil {
		return nil, err
	}

	updatedUser, errUpdate := u.userRepo.UpdateEmail(ctx, sqlc.UpdateEmailParams{
		ID:        int32(id),
		Email:     data.NewEmail,
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if errUpdate != nil {
		return nil, errUpdate
	}

	response := &dto.UpdateUserResponse{
		Message: "Your email has been updated successfully ",
		Data:    sqlc.UpdateUserRow(*updatedUser),
	}

	return response, nil
}

func (u *UserUsecase) UpdatePassword(
	id int,
	data *dto.UpdatePasswordRequest,
) (*dto.UpdatePasswordResponse, error) {
	ctx := context.Background()

	if data.NewPassword != data.ConfirmPassword {
		return nil, fiber.
			NewError(fiber.StatusBadRequest, "Invalid confirm password")
	}

	user, err := u.userRepo.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := u.bcryptService.ComparePassword(
		user.Password,
		data.OldPassword,
	); err != nil {
		return nil, err
	}

	var errHash error
	data.NewPassword, errHash = u.bcryptService.HashPassword(data.NewPassword)
	if errHash != nil {
		return nil, errHash
	}

	if err := u.userRepo.UpdatePassword(ctx, sqlc.UpdatePasswordParams{
		ID:        int32(id),
		Password:  data.NewPassword,
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}); err != nil {
		return nil, err
	}

	response := &dto.UpdatePasswordResponse{
		Message: "Your password has been updated successfully ",
	}

	return response, nil
}

func (u *UserUsecase) DeleteUser(id int) (*dto.DeleteUserResponse, error) {
	ctx := context.Background()

	if _, err := u.userRepo.FindUserByID(ctx, id); err != nil {
		return nil, err
	}

	if err := u.userRepo.DeleteUser(ctx, id); err != nil {
		return nil, err
	}

	response := &dto.DeleteUserResponse{
		Message: "Your user has been deleted successfully ",
	}

	return response, nil
}
