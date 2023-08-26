package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int
	FullName  sql.NullString
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NewUser struct {
	FullName sql.NullString
	Username string
	Email    string
	Password string
}

type CreatedUser struct {
	ID        int
	FullName  sql.NullString
	Username  string
	Email     string
	CreatedAt time.Time
}

type UpdateUser struct {
	FullName sql.NullString
	Username sql.NullString
}

type UpdatedUser struct {
	ID        int
	FullName  sql.NullString
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
