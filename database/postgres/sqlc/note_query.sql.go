// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: note_query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createNote = `-- name: CreateNote :one
INSERT INTO notes (
  user_id, space_id, title, body
) VALUES (
  $1, $2, $3, $4
) RETURNING id, user_id, space_id, title, body, status, created_at
`

type CreateNoteParams struct {
	UserID  int32       `json:"user_id"`
	SpaceID int32       `json:"space_id"`
	Title   string      `json:"title"`
	Body    pgtype.Text `json:"body"`
}

type CreateNoteRow struct {
	ID        int32            `json:"id"`
	UserID    int32            `json:"user_id"`
	SpaceID   int32            `json:"space_id"`
	Title     string           `json:"title"`
	Body      pgtype.Text      `json:"body"`
	Status    Status           `json:"status"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateNote(ctx context.Context, arg CreateNoteParams) (*CreateNoteRow, error) {
	row := q.db.QueryRow(ctx, createNote,
		arg.UserID,
		arg.SpaceID,
		arg.Title,
		arg.Body,
	)
	var i CreateNoteRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.SpaceID,
		&i.Title,
		&i.Body,
		&i.Status,
		&i.CreatedAt,
	)
	return &i, err
}

const deleteNote = `-- name: DeleteNote :exec
DELETE FROM notes WHERE id = $1
`

func (q *Queries) DeleteNote(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteNote, id)
	return err
}

const findAllNotes = `-- name: FindAllNotes :many
SELECT id, user_id, space_id, title, body, status, created_at, updated_at FROM notes WHERE user_id = $1
`

func (q *Queries) FindAllNotes(ctx context.Context, userID int32) ([]*Note, error) {
	rows, err := q.db.Query(ctx, findAllNotes, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Note{}
	for rows.Next() {
		var i Note
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.SpaceID,
			&i.Title,
			&i.Body,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllNotesBySpaceID = `-- name: FindAllNotesBySpaceID :many
SELECT id, user_id, space_id, title, body, status, created_at, updated_at FROM notes WHERE space_id = $1
`

func (q *Queries) FindAllNotesBySpaceID(ctx context.Context, spaceID int32) ([]*Note, error) {
	rows, err := q.db.Query(ctx, findAllNotesBySpaceID, spaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Note{}
	for rows.Next() {
		var i Note
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.SpaceID,
			&i.Title,
			&i.Body,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllNotesBySpaceIDAndStatus = `-- name: FindAllNotesBySpaceIDAndStatus :many
SELECT id, user_id, space_id, title, body, status, created_at, updated_at FROM notes
WHERE space_id = $1 AND status = $2
`

type FindAllNotesBySpaceIDAndStatusParams struct {
	SpaceID int32  `json:"space_id"`
	Status  Status `json:"status"`
}

func (q *Queries) FindAllNotesBySpaceIDAndStatus(ctx context.Context, arg FindAllNotesBySpaceIDAndStatusParams) ([]*Note, error) {
	rows, err := q.db.Query(ctx, findAllNotesBySpaceIDAndStatus, arg.SpaceID, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Note{}
	for rows.Next() {
		var i Note
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.SpaceID,
			&i.Title,
			&i.Body,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllNotesByStatus = `-- name: FindAllNotesByStatus :many
SELECT id, user_id, space_id, title, body, status, created_at, updated_at FROM notes WHERE user_id = $1 AND status = $2
`

type FindAllNotesByStatusParams struct {
	UserID int32  `json:"user_id"`
	Status Status `json:"status"`
}

func (q *Queries) FindAllNotesByStatus(ctx context.Context, arg FindAllNotesByStatusParams) ([]*Note, error) {
	rows, err := q.db.Query(ctx, findAllNotesByStatus, arg.UserID, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Note{}
	for rows.Next() {
		var i Note
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.SpaceID,
			&i.Title,
			&i.Body,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findNoteByID = `-- name: FindNoteByID :one
SELECT id, user_id, space_id, title, body, status, created_at, updated_at FROM notes WHERE id = $1 LIMIT 1
`

func (q *Queries) FindNoteByID(ctx context.Context, id int32) (*Note, error) {
	row := q.db.QueryRow(ctx, findNoteByID, id)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.SpaceID,
		&i.Title,
		&i.Body,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const updateNote = `-- name: UpdateNote :one
UPDATE notes
SET
  title = COALESCE($2, title),
  body = COALESCE($3, body),
  status = COALESCE($4, status),
  space_id = COALESCE($5, space_id),
  updated_at = $6
WHERE id = $1
RETURNING id, space_id, title, body, status, created_at, updated_at
`

type UpdateNoteParams struct {
	ID        int32            `json:"id"`
	Title     pgtype.Text      `json:"title"`
	Body      pgtype.Text      `json:"body"`
	Status    NullStatus       `json:"status"`
	SpaceID   pgtype.Int4      `json:"space_id"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type UpdateNoteRow struct {
	ID        int32            `json:"id"`
	SpaceID   int32            `json:"space_id"`
	Title     string           `json:"title"`
	Body      pgtype.Text      `json:"body"`
	Status    Status           `json:"status"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateNote(ctx context.Context, arg UpdateNoteParams) (*UpdateNoteRow, error) {
	row := q.db.QueryRow(ctx, updateNote,
		arg.ID,
		arg.Title,
		arg.Body,
		arg.Status,
		arg.SpaceID,
		arg.UpdatedAt,
	)
	var i UpdateNoteRow
	err := row.Scan(
		&i.ID,
		&i.SpaceID,
		&i.Title,
		&i.Body,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
