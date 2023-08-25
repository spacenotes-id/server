package usecase

import (
	"context"

	"github.com/tfkhdyt/SpaceNotes/server/internal/application/dto"
	"github.com/tfkhdyt/SpaceNotes/server/internal/domain/repository"
	"github.com/tfkhdyt/SpaceNotes/server/internal/domain/service"
	"github.com/tfkhdyt/SpaceNotes/server/pkg/exception"
)

type UserUsecase struct {
	userRepo       repository.UserRepo    `di.inject:"userRepo"`
	hashingService service.HashingService `di.inject:"hashingService"`
}

func (u *UserUsecase) Register(
	newUser *dto.RegisterRequest,
) (*dto.RegisterResponse, error) {
	ctx := context.Background()

	if _, err := u.userRepo.FindUserByUsername(
		ctx,
		newUser.Username,
	); err == nil {
		return nil, exception.NewHTTPError(400, "username has been used")
	}

	if _, err := u.userRepo.FindUserByEmail(ctx, newUser.Email); err == nil {
		return nil, exception.NewHTTPError(400, "email has been used")
	}

	var errHash error
	newUser.Password, errHash = u.hashingService.HashPassword(newUser.Password)
	if errHash != nil {
		return nil, errHash
	}

	registeredUser, errRegister := u.userRepo.CreateUser(ctx, newUser)
	if errRegister != nil {
		return nil, errRegister
	}

	response := &dto.RegisterResponse{
		Message: "Your account has been created successfully",
		Data:    *registeredUser,
	}

	return response, nil
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

	updatedUser, err := u.userRepo.UpdateUser(ctx, id, data)
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
		Message: "Your email has been updated successfully ",
		Data:    *updatedUser,
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
