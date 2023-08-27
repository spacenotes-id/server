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
)
