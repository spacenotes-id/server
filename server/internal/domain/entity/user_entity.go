package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int            `json:"id"`
	FullName  sql.NullString `json:"full_name"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type NewUser struct {
	FullName string `json:"full_name" valid:"stringlength(2|70)~Full name length should be at least 2 - 70 characters"`
	Username string `json:"username" valid:"required~Username is required,stringlength(3|16)~Username length should be at least 3 - 16 characters"`
	Email    string `json:"email" valid:"required~Email is required,email~Invalid email"`
	Password string `json:"password" valid:"required~Password is required"`
}

type CreatedUser struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateUser struct {
	FullName string `json:"full_name" valid:"stringlength(2|70)~Full name length should be at least 2 - 70 characters"`
	Username string `json:"username" valid:"stringlength(3|16)~Username length should be at least 3 - 16 characters"`
}

type UpdatedUser struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
