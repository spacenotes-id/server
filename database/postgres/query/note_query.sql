-- name: CreateNote :one
INSERT INTO notes (
  user_id, space_id, title, body
) VALUES (
  $1, $2, $3, $4
) RETURNING id, user_id, space_id, title, body, status, created_at;

-- name: FindAllNotes :many
SELECT * FROM notes WHERE user_id = $1;

-- name: FindAllNotesByStatus :many
SELECT * FROM notes WHERE user_id = $1 AND status = $2;

-- name: FindAllNotesBySpaceID :many
SELECT * FROM notes WHERE space_id = $1;

-- name: FindAllNotesBySpaceIDAndStatus :many
SELECT * FROM notes
WHERE space_id = $1 AND status = $2;

-- name: FindNoteByID :one
SELECT * FROM notes WHERE id = $1 LIMIT 1;

-- name: UpdateNote :one
UPDATE notes
SET
  title = COALESCE(sqlc.narg('title'), title),
  body = COALESCE(sqlc.narg('body'), body),
  status = COALESCE(sqlc.narg('status'), status),
  space_id = COALESCE(sqlc.narg('space_id'), space_id),
  updated_at = sqlc.arg('updated_at')
WHERE id = $1
RETURNING id, space_id, title, body, status, created_at, updated_at;

-- name: DeleteNote :exec
DELETE FROM notes WHERE id = $1;
