-- name: CreateUser :one
INSERT INTO "user" (
  full_name, username, email, password
) VALUES (
  $1, $2, $3, $4
) RETURNING id, full_name, username, email, created_at;

-- name: FindUserByID :one
SELECT * FROM "user" WHERE id = $1 LIMIT 1;

-- name: FindUserByUsername :one
SELECT * FROM "user" WHERE username = $1 LIMIT 1;

-- name: FindUserByEmail :one
SELECT * FROM "user" WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE "user"
SET
  full_name = coalesce(sqlc.narg('fullName'), full_name),
  username = coalesce(sqlc.narg('username'), username)
WHERE id = sqlc.arg('id')
RETURNING id, full_name, username, email, created_at, updated_at;

-- name: UpdateEmail :one
UPDATE "user"
SET email = $2
WHERE id = $1
RETURNING id, full_name, username, email, created_at, updated_at;

-- name: UpdatePassword :exec
UPDATE "user"
SET password = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM "user" WHERE id = $1;