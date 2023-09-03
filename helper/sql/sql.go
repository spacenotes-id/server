package sql

import (
	"database/sql"

	"github.com/spacenotes-id/server/database/postgres/sqlc"
)

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func NewNullBool(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{}
	}
	return sql.NullBool{
		Bool:  *b,
		Valid: true,
	}
}

func NewNullStatus(s string) sqlc.NullStatus {
	if len(s) == 0 {
		return sqlc.NullStatus{}
	}
	return sqlc.NullStatus{
		Status: sqlc.Status(s),
		Valid:  true,
	}
}

func NewNullInt(i int) sql.NullInt32 {
	if i == 0 {
		return sql.NullInt32{}
	}

	return sql.NullInt32{
		Int32: int32(i),
		Valid: true,
	}
}
