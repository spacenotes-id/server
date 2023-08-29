-- name: CreateSpace :one
INSERT INTO spaces (name, emoji, is_locked, user_id) VALUES ($1, $2, $3, $4)
RETURNING id, user_id, name, emoji, is_locked, created_at;

-- name: FindAllSpacesByUserID :many
SELECT id, name, emoji, is_locked, created_at, updated_at FROM spaces 
WHERE user_id = $1;

-- name: FindSpaceByID :one
SELECT * FROM spaces WHERE user_id = $1 AND id = $2;

-- name: FindSpaceByName :one
SELECT * FROM spaces WHERE name = $1;

-- name: UpdateSpace :one
UPDATE spaces
SET
  name = COALESCE(sqlc.narg('name'), name),
  emoji = COALESCE(sqlc.narg('emoji'), emoji),
  is_locked = COALESCE(sqlc.narg('is_locked'), is_locked),
  updated_at = sqlc.arg('updated_at')
WHERE user_id = $1 AND id = $2 
RETURNING id, name, emoji, is_locked, created_at, updated_at;

-- name: DeleteSpace :exec
DELETE FROM spaces WHERE user_id = $1 AND id = $2;
