-- name: AddToken :one
INSERT INTO refresh_tokens (token) VALUES ($1) RETURNING token;

-- name: FindToken :one
SELECT * FROM refresh_tokens WHERE token = $1;

-- name: DeleteToken :exec
DELETE FROM refresh_tokens WHERE token = $1;
