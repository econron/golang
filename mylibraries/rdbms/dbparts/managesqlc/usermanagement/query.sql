-- name: CreateUser :exec
INSERT INTO users (username, email, password) 
VALUES ($1, $2, $3);

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users 
SET email = $2, password = $3 
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
