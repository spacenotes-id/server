package dto

import "time"

type (
	RegisterRequest struct {
		FullName string `json:"full_name" valid:"stringlength(2|70)~Full name length should be at least 2 - 70 characters"`
		Username string `json:"username" valid:"required~Username is required,stringlength(3|16)~Username length should be at least 3 - 16 characters"`
		Email    string `json:"email" valid:"required~Email is required,email~Invalid email"`
		Password string `json:"password" valid:"required~Password is required"`
	}
	RegisterResponseData struct {
		ID        int       `json:"id"`
		FullName  string    `json:"full_name"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
	}
	RegisterResponse struct {
		Message string               `json:"message"`
		Data    RegisterResponseData `json:"data"`
	}

	FindUserByIDResponseData struct {
		ID        int       `json:"id"`
		FullName  string    `json:"full_name"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	FindUserByIDResponse struct {
		Data FindUserByIDResponseData `json:"data"`
	}

	UpdateUserRequest struct {
		FullName string `json:"full_name" valid:"stringlength(2|70)~Full name length should be at least 2 - 70 characters"`
		Username string `json:"username" valid:"stringlength(3|16)~Username length should be at least 3 - 16 characters"`
	}
	UpdateUserResponseData struct {
		ID        int       `json:"id"`
		FullName  string    `json:"full_name"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	UpdateUserResponse struct {
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
