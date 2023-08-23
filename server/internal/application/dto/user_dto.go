package dto

import "github.com/tfkhdyt/SpaceNotes/server/internal/domain/entity"

type (
	RegisterRequest      = entity.NewUser
	RegisterResponseData = entity.CreatedUser
	RegisterResponse     struct {
		Message string               `json:"message"`
		Data    RegisterResponseData `json:"data"`
	}

	FindUserByIDResponseData = entity.User
	FindUserByIDResponse     struct {
		Data FindUserByIDResponseData `json:"data"`
	}

	UpdateUserRequest      = entity.UpdateUser
	UpdateUserResponseData = entity.UpdatedUser
	UpdateUserResponse     struct {
		Message string                 `json:"message"`
		Data    UpdateUserResponseData `json:"data"`
	}

	UpdateEmailRequest struct {
		NewEmail string `json:"new_email"`
		Password string `json:"password"`
	}

	UpdatePasswordRequest struct {
		OldPassword     string `json:"old_password"`
		NewPassword     string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}
	UpdatePasswordResponse struct {
		Message string `json:"message"`
	}

	DeleteUserResponse = UpdatePasswordResponse
)
