-- name: AddToken :exec
INSERT INTO refresh_tokens (token) VALUES ($1);

-- name: FindToken :one
SELECT * FROM refresh_tokens WHERE token = $1 LIMIT 1;

-- name: DeleteToken :exec
DELETE FROM refresh_tokens WHERE token = $1;
