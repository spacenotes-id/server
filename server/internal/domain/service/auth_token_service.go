package service

type AuthTokenService interface {
	CreateAccessToken(id int) (string, error)
	CreateRefreshToken(id int) (string, error)
	ParseRefreshToken(tokenString string) (int, error)
}
