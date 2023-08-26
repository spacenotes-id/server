package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tfkhdyt/SpaceNotes/server/config"
	"github.com/tfkhdyt/SpaceNotes/server/pkg/exception"
)

type JwtService struct{}

func (j *JwtService) CreateAccessToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":    "spacenotes-server",
		"userId": id,
		"exp":    time.Now().Add(15 * time.Minute).Unix(),
	})
	if token == nil {
		return "", exception.
			NewHTTPError(500, "failed to to create new access token")
	}

	signedString, err := token.SignedString([]byte(config.JwtAccessTokenKey))
	if err != nil {
		return "", exception.NewHTTPError(500, "failed to sign access token")
	}

	return signedString, nil
}

func (j *JwtService) CreateRefreshToken(
	id int,
) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":    "spacenotes-server",
		"userId": id,
		"exp":    time.Now().Add(720 * time.Hour).Unix(),
	})
	if token == nil {
		return "", exception.
			NewHTTPError(500, "failed to to create new refresh token")
	}

	signedString, err := token.SignedString([]byte(config.JwtRefreshTokenKey))
	if err != nil {
		return "", exception.NewHTTPError(500, "failed to sign refresh token")
	}

	return signedString, nil
}

func (j *JwtService) ParseRefreshToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, exception.NewHTTPError(
				400,
				fmt.Sprintf(
					"unexpected signing method: %v",
					token.Header["alg"],
				),
			)
		}

		return []byte(config.JwtRefreshTokenKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, exception.NewHTTPError(400, "invalid token")
	}

	userId, okUserId := claims["userId"].(float64)
	if !okUserId {
		return 0, exception.
			NewHTTPError(400, "failed to parse user id from claims")
	}

	return int(userId), nil
}
