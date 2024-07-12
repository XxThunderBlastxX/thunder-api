// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package projectDb

import (
	"context"

	"github.com/lib/pq"
)

const createProject = `-- name: CreateProject :one
INSERT INTO projects (id, name, description, link, stacks)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, description, link, stacks, created_at, updated_at
`

type CreateProjectParams struct {
	ID          int32
	Name        string
	Description string
	Link        string
	Stacks      []string
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, createProject,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Link,
		pq.Array(arg.Stacks),
	)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Link,
		pq.Array(&i.Stacks),
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProjectById = `-- name: DeleteProjectById :exec
DELETE FROM projects
WHERE id = $1
`

func (q *Queries) DeleteProjectById(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteProjectById, id)
	return err
}

const deleteProjectByName = `-- name: DeleteProjectByName :exec
DELETE FROM projects
WHERE name = $1
`

func (q *Queries) DeleteProjectByName(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deleteProjectByName, name)
	return err
}

const listProjects = `-- name: ListProjects :many
SELECT id, name, description, link, stacks, created_at, updated_at FROM projects
`

func (q *Queries) ListProjects(ctx context.Context) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, listProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Link,
			pq.Array(&i.Stacks),
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
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