package usecase

import (
	"context"

	"github.com/tfkhdyt/SpaceNotes/server/internal/application/dto"
	"github.com/tfkhdyt/SpaceNotes/server/internal/domain/entity"
	"github.com/tfkhdyt/SpaceNotes/server/internal/domain/repository"
	"github.com/tfkhdyt/SpaceNotes/server/internal/domain/service"
	"github.com/tfkhdyt/SpaceNotes/server/pkg/exception"
	"github.com/tfkhdyt/SpaceNotes/server/pkg/sql"
)

type UserUsecase struct {
	userRepo       repository.UserRepo    `di.inject:"userRepo"`
	hashingService service.HashingService `di.inject:"hashingService"`
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
			ID:        user.ID,
			FullName:  user.FullName.String,
			Username:  user.Username,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
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
		if _, err := u.userRepo.FindUserByUsername(ctx, data.Username); err != nil {
			return nil, err
		}
	}

	updatedUser, err := u.userRepo.UpdateUser(ctx, id, &entity.UpdateUser{
		FullName: sql.NewNullString(data.FullName),
		Username: sql.NewNullString(data.Username),
	})
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateUserResponse{
		Message: "Your account data has been updated successfully ",
		Data: dto.UpdateUserResponseData{
			ID:        updatedUser.ID,
			FullName:  updatedUser.FullName.String,
			Username:  updatedUser.Username,
			Email:     updatedUser.Email,
			CreatedAt: updatedUser.CreatedAt,
			UpdatedAt: updatedUser.UpdatedAt,
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

	if err := u.hashingService.ComparePassword(
		user.Password,
		data.Password,
	); err != nil {
		return nil, err
	}

	updatedUser, errUpdate := u.userRepo.UpdateEmail(ctx, id, data.NewEmail)
	if errUpdate != nil {
		return nil, errUpdate
	}

	response := &dto.UpdateUserResponse{
		Message: "Your account data has been updated successfully ",
		Data: dto.UpdateUserResponseData{
			ID:        updatedUser.ID,
			FullName:  updatedUser.FullName.String,
			Username:  updatedUser.Username,
			Email:     updatedUser.Email,
			CreatedAt: updatedUser.CreatedAt,
			UpdatedAt: updatedUser.UpdatedAt,
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
		return nil, exception.NewHTTPError(400, "Invalid confirm password")
	}

	user, err := u.userRepo.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := u.hashingService.ComparePassword(
		user.Password,
		data.OldPassword,
	); err != nil {
		return nil, err
	}

	var errHash error
	data.NewPassword, errHash = u.hashingService.HashPassword(data.NewPassword)
	if errHash != nil {
		return nil, errHash
	}

	if err := u.userRepo.UpdatePassword(ctx, id, data.NewPassword); err != nil {
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
