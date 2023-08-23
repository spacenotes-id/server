package repository

import "github.com/tfkhdyt/SpaceNotes/server/internal/domain/entity"

type UserRepo interface {
	CreateUser(user *entity.NewUser) (*entity.CreatedUser, error)

	FindUserByID(id int) (*entity.User, error)
	FindUserByUsername(username string) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)

	UpdateUser(id int, data *entity.UpdateUser) (*entity.UpdatedUser, error)
	UpdateEmail(id int, email string) (*entity.UpdatedUser, error)
	UpdatePassword(id int, password string) error

	DeleteUser(id int) error
}
