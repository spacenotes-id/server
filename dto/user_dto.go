package dto

import (
	"time"
)

type (
	FindUserByIDData struct {
		ID        int32     `json:"id"         example:"1"`
		FullName  string    `json:"full_name"  example:"Taufik Hidayat"`
		Username  string    `json:"username"   example:"tfkhdyt"`
		Email     string    `json:"email"      example:"me@tfkhdyt.my.id"`
		CreatedAt time.Time `json:"created_at" example:"2023-09-22T04:11:26.597Z"`
		UpdatedAt time.Time `json:"updated_at" example:"2023-09-22T04:11:26.597Z"`
	}
	FindUserByIDResponse struct {
		Data FindUserByIDData `json:"data"`
	}

	UpdateUserRequest struct {
		FullName string `json:"full_name" valid:"stringlength(2|70)~Full name length should be at least 2 - 70 characters" validate:"optional" minLength:"2" maxLength:"70" example:"Fauzi Fathirohman"`
		Username string `json:"username"  valid:"stringlength(3|16)~Username length should be at least 3 - 16 characters"  validate:"optional" minLength:"3" maxLength:"16" example:"fauzi123"`
	}
	UpdateUserResponse struct {
		Message string           `json:"message" example:"Your account data has been updated successfully "`
		Data    FindUserByIDData `json:"data"`
	}

	UpdateEmailRequest struct {
		NewEmail string `json:"new_email" valid:"required~New email is required,email~Invalid email" validate:"required" example:"tfkhdyt@proton.me"`
		Password string `json:"password"  valid:"required~Password is required"                      validate:"required" example:"bruh1234"`
	}

	UpdatePasswordRequest struct {
		OldPassword     string `json:"old_password"     valid:"required~Old password is required"`
		NewPassword     string `json:"new_password"     valid:"required~New password is required"`
		ConfirmPassword string `json:"confirm_password" valid:"required~Confirm password is required"`
	}
	UpdatePasswordResponse struct {
		Message string `json:"message"`
	}

	DeleteUserResponse = UpdatePasswordResponse
)
