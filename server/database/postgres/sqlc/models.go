// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        int32            `json:"id"`
	FullName  pgtype.Text      `json:"full_name"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}
