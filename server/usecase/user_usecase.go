package usecase

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tfkhdyt/SpaceNotes/server/database/postgres/sqlc"
	"github.com/tfkhdyt/SpaceNotes/server/dto"
	"github.com/tfkhdyt/SpaceNotes/server/helper/sql"
	"github.com/tfkhdyt/SpaceNotes/server/repository/postgres"
	"github.com/tfkhdyt/SpaceNotes/server/service"
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
		Data: dto.FindUserByIDResponseData{
			ID:        int(user.ID),
			FullName:  user.FullName.String,
			Username:  user.Username,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt.Time,
			UpdatedAt: user.UpdatedAt.Time,
		},
	}

	return response, nil
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
		if _, err := u.userRepo.FindUserByUsername(
			ctx,
			data.Username,
		); err != nil {
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
		Data: dto.UpdateUserResponseData{
			ID:        int(updatedUser.ID),
			FullName:  updatedUser.FullName.String,
			Username:  updatedUser.Username,
			Email:     updatedUser.Email,
			CreatedAt: updatedUser.CreatedAt.Time,
			UpdatedAt: updatedUser.UpdatedAt.Time,
		},
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

	if _, err := u.userRepo.FindUserByEmail(ctx, data.NewEmail); err != nil {
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
		Message: "Your account data has been updated successfully ",
		Data: dto.UpdateUserResponseData{
			ID:        int(updatedUser.ID),
			FullName:  updatedUser.FullName.String,
			Username:  updatedUser.Username,
			Email:     updatedUser.Email,
			CreatedAt: updatedUser.CreatedAt.Time,
			UpdatedAt: updatedUser.UpdatedAt.Time,
		},
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
