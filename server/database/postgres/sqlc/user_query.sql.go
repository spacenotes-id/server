// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user_query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  full_name, username, email, password
) VALUES (
  $1, $2, $3, $4
) RETURNING id, full_name, username, email, created_at
`

type CreateUserParams struct {
	FullName pgtype.Text `json:"full_name"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
}

type CreateUserRow struct {
	ID        int32            `json:"id"`
	FullName  pgtype.Text      `json:"full_name"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (*CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.FullName,
		arg.Username,
		arg.Email,
		arg.Password,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Username,
		&i.Email,
		&i.CreatedAt,
	)
	return &i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT id, full_name, username, email, password, created_at, updated_at FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	row := q.db.QueryRow(ctx, findUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const findUserByID = `-- name: FindUserByID :one
SELECT id, full_name, username, email, password, created_at, updated_at FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) FindUserByID(ctx context.Context, id int32) (*User, error) {
	row := q.db.QueryRow(ctx, findUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const findUserByUsername = `-- name: FindUserByUsername :one
SELECT id, full_name, username, email, password, created_at, updated_at FROM users WHERE username = $1 LIMIT 1
`

func (q *Queries) FindUserByUsername(ctx context.Context, username string) (*User, error) {
	row := q.db.QueryRow(ctx, findUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const updateEmail = `-- name: UpdateEmail :one
UPDATE users
SET 
  email = $2,
  updated_at = $3
WHERE id = $1
RETURNING id, full_name, username, email, created_at, updated_at
`

type UpdateEmailParams struct {
	ID        int32            `json:"id"`
	Email     string           `json:"email"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type UpdateEmailRow struct {
	ID        int32            `json:"id"`
	FullName  pgtype.Text      `json:"full_name"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateEmail(ctx context.Context, arg UpdateEmailParams) (*UpdateEmailRow, error) {
	row := q.db.QueryRow(ctx, updateEmail, arg.ID, arg.Email, arg.UpdatedAt)
	var i UpdateEmailRow
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Username,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE users
SET 
  password = $2,
  updated_at = $3
WHERE id = $1
`

type UpdatePasswordParams struct {
	ID        int32            `json:"id"`
	Password  string           `json:"password"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.Exec(ctx, updatePassword, arg.ID, arg.Password, arg.UpdatedAt)
	return err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
  full_name = COALESCE($1, full_name),
  username = COALESCE($2, username),
  updated_at = $3
WHERE id = $4
RETURNING id, full_name, username, email, created_at, updated_at
`

type UpdateUserParams struct {
	FullName  pgtype.Text      `json:"full_name"`
	Username  pgtype.Text      `json:"username"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	ID        int32            `json:"id"`
}

type UpdateUserRow struct {
	ID        int32            `json:"id"`
	FullName  pgtype.Text      `json:"full_name"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (*UpdateUserRow, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.FullName,
		arg.Username,
		arg.UpdatedAt,
		arg.ID,
	)
	var i UpdateUserRow
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Username,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
