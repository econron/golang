-- name: FindUserById :one
SELECT * FROM users WHERE id = ?;

-- name: CreateUser :exec
INSERT INTO users (name, email)
VALUES (?, ?);

-- name: UpdateUser :exec
UPDATE users
SET name = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
