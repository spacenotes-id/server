package entity

import "time"

type User struct {
	ID        int
	FullName  *string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NewUser struct {
	FullName *string
	Username string
	Email    string
	Password string
}

type CreatedUser struct {
	ID        int
	FullName  *string
	Username  string
	Email     string
	CreatedAt time.Time
}

type UpdateUser struct {
	FullName *string
	Username *string
}

type UpdatedUser struct {
	ID        int
	FullName  *string
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
