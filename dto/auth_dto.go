package dto

import (
	"time"
)

type (
	RegisterRequest struct {
		FullName string `json:"full_name" valid:"stringlength(2|70)~Full name length should be at least 2 - 70 characters" validate:"optional" minLength:"2" maxLength:"70" example:"Taufik Hidayat"`
		Username string `json:"username" valid:"required~Username is required,stringlength(3|16)~Username length should be at least 3 - 16 characters" validate:"required" minLength:"3" maxLength:"16" example:"tfkhdyt"`
		Email    string `json:"email" valid:"required~Email is required,email~Invalid email" validate:"required" example:"me@tfkhdyt.my.id"`
		Password string `json:"password" valid:"required~Password is required,minstringlength(8)~Your password should be more than 8 characters" validate:"required" example:"bruh1234" minLength:"8"`
	}
	RegisterResponseData struct {
		ID        int32     `json:"id" example:"69"`
		FullName  string    `json:"full_name" example:"Taufik Hidayat"`
		Username  string    `json:"username" example:"tfkhdyt"`
		Email     string    `json:"email" example:"me@tfkhdyt.my.id"`
		CreatedAt time.Time `json:"created_at" example:"2023-09-04T21:00:43.775157Z"`
	}
	RegisterResponse struct {
		Message string               `json:"message" example:"Your account has been created successfully"`
		Data    RegisterResponseData `json:"data"`
	}

	LoginRequest struct {
		Email    string `json:"email" valid:"required~Email is required,email~Invalid email"`
		Password string `json:"password" valid:"required~Password is required"`
	}
	LoginResponseData struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	LoginResponse struct {
		Message string            `json:"message"`
		Data    LoginResponseData `json:"data"`
	}

	LogoutRequest struct {
		RefreshToken string `json:"refresh_token" valid:"required~Refresh token is required"`
	}
	LogoutResponse struct {
		Message string `json:"message"`
	}

	RefreshRequest      = LogoutRequest
	RefreshResponseData struct {
		AccessToken string `json:"access_token"`
	}
	RefreshResponse struct {
		Message string              `json:"message"`
		Data    RefreshResponseData `json:"data"`
	}
)
