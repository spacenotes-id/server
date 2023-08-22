package service

type HashingService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) error
}
