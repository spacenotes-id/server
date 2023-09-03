package service

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spacenotes-id/server/config"
)

type JwtService struct{}

func (j *JwtService) CreateAccessToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":     "spacenotes-server",
		"user_id": id,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	})
	if token == nil {
		return "", fiber.NewError(
			fiber.StatusInternalServerError,
			"Failed to create new access token",
		)
	}

	signedString, err := token.SignedString([]byte(config.JwtAccessTokenKey))
	if err != nil {
		log.Error(err)
		return "", fiber.
			NewError(fiber.StatusInternalServerError, "Failed to sign access token")
	}

	return signedString, nil
}

func (j *JwtService) CreateRefreshToken(
	id int,
) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":     "spacenotes-server",
		"user_id": id,
		"exp":     time.Now().Add(720 * time.Hour).Unix(),
	})
	if token == nil {
		return "", fiber.NewError(
			fiber.StatusInternalServerError,
			"Failed to create new refresh token",
		)
	}

	signedString, err := token.SignedString([]byte(config.JwtRefreshTokenKey))
	if err != nil {
		log.Error(err)
		return "", fiber.
			NewError(fiber.StatusInternalServerError, "Failed to sign refresh token")
	}

	return signedString, nil
}

func (j *JwtService) ParseRefreshToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf(
				"unexpected signing method: %v",
				token.Header["alg"],
			))
		}

		return []byte(config.JwtRefreshTokenKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Invalid token")
	}

	userId, okUserId := claims["user_id"].(float64)
	if !okUserId {
		return 0, fiber.
			NewError(fiber.StatusBadRequest, "Failed to parse user id from claims")
	}

	return int(userId), nil
}
