package dto

type (
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
