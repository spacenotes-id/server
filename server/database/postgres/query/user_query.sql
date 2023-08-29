-- name: CreateUser :one
INSERT INTO users (
  full_name, username, email, password
) VALUES (
  $1, $2, $3, $4
) RETURNING id, full_name, username, email, created_at;

-- name: FindUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: FindUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  full_name = COALESCE(sqlc.narg('full_name'), full_name),
  username = COALESCE(sqlc.narg('username'), username),
  updated_at = sqlc.arg('updated_at')
WHERE id = sqlc.arg('id')
RETURNING id, full_name, username, email, created_at, updated_at;

-- name: UpdateEmail :one
UPDATE users
SET 
  email = $2,
  updated_at = $3
WHERE id = $1
RETURNING id, full_name, username, email, created_at, updated_at;

-- name: UpdatePassword :exec
UPDATE users
SET 
  password = $2,
  updated_at = $3
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
