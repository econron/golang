-- name: CreateAdmin :exec
INSERT INTO admin (username, email, password, role) 
VALUES ($1, $2, $3, $4);

-- name: GetAdminById :one
SELECT * FROM admin WHERE id = $1;

-- name: UpdateAdmin :exec
UPDATE admin 
SET email = $2, password = $3, role = $4 
WHERE id = $1;

-- name: DeleteAdmin :exec
DELETE FROM admin WHERE id = $1;
