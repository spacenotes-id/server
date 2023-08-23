package usecase

import (
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
	if _, err := u.userRepo.FindUserByUsername(newUser.Username); err == nil {
		return nil, err
	}

	if _, err := u.userRepo.FindUserByEmail(newUser.Email); err == nil {
		return nil, err
	}

	var errHash error
	newUser.Password, errHash = u.hashingService.HashPassword(newUser.Password)
	if errHash != nil {
		return nil, errHash
	}

	registeredUser, errRegister := u.userRepo.CreateUser(newUser)
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
	user, err := u.userRepo.FindUserByID(id)
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
	if _, err := u.userRepo.FindUserByID(id); err != nil {
		return nil, err
	}

	if data.Username != nil {
		if _, err := u.userRepo.FindUserByUsername(*data.Username); err != nil {
			return nil, err
		}
	}

	updatedUser, err := u.userRepo.UpdateUser(id, data)
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
	user, err := u.userRepo.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	if _, err := u.userRepo.FindUserByEmail(data.NewEmail); err != nil {
		return nil, err
	}

	if err := u.hashingService.ComparePassword(
		user.Password,
		data.Password,
	); err != nil {
		return nil, err
	}

	updatedUser, errUpdate := u.userRepo.UpdateEmail(id, data.NewEmail)
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
	if data.NewPassword != data.ConfirmPassword {
		return nil, exception.NewHTTPError(400, "Invalid confirm password")
	}

	user, err := u.userRepo.FindUserByID(id)
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

	if err := u.userRepo.UpdatePassword(id, data.NewPassword); err != nil {
		return nil, err
	}

	response := &dto.UpdatePasswordResponse{
		Message: "Your password has been updated successfully ",
	}

	return response, nil
}

func (u *UserUsecase) DeleteUser(id int) (*dto.DeleteUserResponse, error) {
	if _, err := u.userRepo.FindUserByID(id); err != nil {
		return nil, err
	}

	if err := u.userRepo.DeleteUser(id); err != nil {
		return nil, err
	}

	response := &dto.DeleteUserResponse{
		Message: "Your user has been deleted successfully ",
	}

	return response, nil
}
