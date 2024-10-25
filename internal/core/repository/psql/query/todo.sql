-- List all todos :many
SELECT *
FROM "todos";

-- Count all todos :one
SELECT count(*)
FROM "todos";


-- Create a new to do :exec
INSERT INTO "todos" (title, description)
VALUES ($1, $2);


-- Get a to do by id :one
SELECT *
FROM "todos"
WHERE id = $1;


-- Update a to do by id :exec
UPDATE "todos"
SET title       = $1,
    description = $2,
    completed   = $3,
    updated_at  = NOW()
WHERE id = $4;