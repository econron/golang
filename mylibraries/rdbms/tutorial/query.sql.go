// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package tutorial

import (
	"context"
	"database/sql"
)

const createAutor = `-- name: CreateAutor :execresult
INSERT INTO authors (
    name, bio
) VALUES (
    ?, ?
)
`

type CreateAutorParams struct {
	Name string
	Bio  sql.NullString
}

func (q *Queries) CreateAutor(ctx context.Context, arg CreateAutorParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAutor, arg.Name, arg.Bio)
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = ?
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, bio FROM authors
WHERE id = ? LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const listAuthros = `-- name: ListAuthros :many
SELECT id, name, bio FROM authors
ORDER BY name
`

func (q *Queries) ListAuthros(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthros)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
