// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package dbaccess

import (
	"context"
)

const createAdmin = `-- name: CreateAdmin :exec
INSERT INTO admin (username, email, password, role) 
VALUES ($1, $2, $3, $4)
`

func (q *Queries) CreateAdmin(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createAdmin)
	return err
}

const deleteAdmin = `-- name: DeleteAdmin :exec
DELETE FROM admin WHERE id = $1
`

func (q *Queries) DeleteAdmin(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAdmin)
	return err
}

const getAdminById = `-- name: GetAdminById :one
SELECT id, username, email, password, role, created_at, updated_at FROM admin WHERE id = $1
`

func (q *Queries) GetAdminById(ctx context.Context) (Admin, error) {
	row := q.db.QueryRowContext(ctx, getAdminById)
	var i Admin
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateAdmin = `-- name: UpdateAdmin :exec
UPDATE admin 
SET email = $2, password = $3, role = $4 
WHERE id = $1
`

func (q *Queries) UpdateAdmin(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updateAdmin)
	return err
}
