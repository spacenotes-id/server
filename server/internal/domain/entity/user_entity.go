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
	FullName *string `json:"full_name" valid:"stringlength(2|70)~Full name length should be at least 2 - 70 characters"`
	Username string  `json:"username" valid:"required~Username is required,stringlength(3|16)~Username length should be at least 3 - 16 characters"`
	Email    string  `json:"email" valid:"required~Email is required,email~Invalid email"`
	Password string  `json:"password" valid:"required~Password is required"`
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
