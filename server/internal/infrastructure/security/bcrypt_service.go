package security

import (
	"log"

	"github.com/tfkhdyt/SpaceNotes/server/pkg/exception"
	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct{}

func (b *BcryptService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Println("ERROR(HashPassword):", err)
		return "", exception.NewHTTPError(500, "Failed to hash password")
	}

	return string(hashedPassword), nil
}

func (b *BcryptService) ComparePassword(
	hashedPassword string,
	password string,
) error {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	); err != nil {
		return exception.NewHTTPError(400, "Password is invalid")
	}

	return nil
}
