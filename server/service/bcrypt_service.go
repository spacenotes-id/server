package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct{}

func (b *BcryptService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Error(err)
		return "", fiber.
			NewError(fiber.StatusInternalServerError, "Failed to hash password")
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
		return fiber.NewError(fiber.StatusBadRequest, "Password is invalid")
	}

	return nil
}
