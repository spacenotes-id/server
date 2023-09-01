-- name: CreateNote :one
INSERT INTO notes (
  user_id, space_id, title, body
) VALUES (
  $1, $2, $3, $4
) RETURNING id, user_id, space_id, title, body, status, created_at;

-- name: FindAllNotes :many
SELECT id, space_id, title, body, status, created_at, updated_at
FROM notes
WHERE user_id = $1 AND status < 'archived'::status;

-- name: FindAllNotesBySpaceID :many
SELECT id, title, body, status, created_at, updated_at
FROM notes
WHERE space_id = $1 AND status < 'archived'::status;

-- name: FindAllTrashedNotes :many
SELECT id, space_id, title, body, created_at, updated_at
FROM notes
WHERE user_id = $1 AND status = 'trashed'::status;

-- name: FindAllFavoriteNotes :many
SELECT id, space_id, title, body, created_at, updated_at
FROM notes
WHERE user_id = $1 AND status = 'favorite'::status;

-- name: FindAllArchivedNotes :many
SELECT id, space_id, title, body, created_at, updated_at
FROM notes
WHERE user_id = $1 AND status = 'archived'::status;

-- name: FindNoteByID :one
SELECT * FROM notes WHERE id = $1;

-- name: UpdateNote :one
UPDATE notes
SET
  title = COALESCE(sqlc.narg('title'), title),
  body = COALESCE(sqlc.narg('body'), body),
  status = COALESCE(sqlc.narg('status'), status),
  updated_at = sqlc.arg('updated_at')
WHERE id = $1
RETURNING id, space_id, title, body, status, created_at, updated_at;

-- name: DeleteNote :exec
DELETE FROM notes WHERE id = $1;
