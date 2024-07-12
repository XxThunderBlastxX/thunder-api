-- name: ListProjects :many
SELECT * FROM projects;

-- name: CreateProject :one
INSERT INTO projects (id, name, description, link, stacks)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteProjectById :exec
DELETE FROM projects
WHERE id = $1;

-- name: DeleteProjectByName :exec
DELETE FROM projects
WHERE name = $1;