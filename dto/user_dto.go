package dto

import (
	"github.com/spacenotes-id/server/database/postgres/sqlc"
)

type (
	FindUserByIDResponse struct {
		Data sqlc.User `json:"data"`
	}

	UpdateUserRequest struct {
		FullName string `json:"full_name" valid:"stringlength(2|70)~Full name length should be at least 2 - 70 characters"`
		Username string `json:"username" valid:"stringlength(3|16)~Username length should be at least 3 - 16 characters"`
	}
	UpdateUserResponse struct {
		Message string             `json:"message"`
		Data    sqlc.UpdateUserRow `json:"data"`
	}

	UpdateEmailRequest struct {
		NewEmail string `json:"new_email" valid:"required~New email is required,email~Invalid email"`
		Password string `json:"password" valid:"required~Password is required"`
	}

	UpdatePasswordRequest struct {
		OldPassword     string `json:"old_password" valid:"required~Old password is required"`
		NewPassword     string `json:"new_password" valid:"required~New password is required"`
		ConfirmPassword string `json:"confirm_password" valid:"required~Confirm password is required"`
	}
	UpdatePasswordResponse struct {
		Message string `json:"message"`
	}

	DeleteUserResponse = UpdatePasswordResponse
)
