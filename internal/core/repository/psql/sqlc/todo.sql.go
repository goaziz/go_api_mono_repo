package sqlc

import "context"

const countTodo = `-- name: CountTodo :one 
SELECT count(*) FROM "todo"`

func (q *Queries) CountTodo(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countTodo)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createTodo = `-- name: CreateTodo :exec
INSERT INTO "todo" (title, description) VALUES ($1, $2)
RETURNING id, title, description, created_at, updated_at;`

type CreateTodoParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	var todo Todo
	err := q.db.QueryRow(ctx, createTodo, arg.Title, arg.Description).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.CreatedAt,
		&todo.UpdateAt,
	)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

const getTodoById = `-- name: GetTodoById :one
SELECT id, title, description, completed, created_at, updated_at FROM "todo"
WHERE id = $1;`

func (q *Queries) GetTodoById(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRow(ctx, getTodoById, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, title, description, completed, created_at FROM "todo"`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.Query(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Completed,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE "todo"
SET title = $1, description = $2, completed = $3, updated_at = now()
WHERE id = $4
RETURNING id, title, description, completed, created_at, updated_at`

type UpdateTodoParams struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	var todo Todo
	err := q.db.QueryRow(ctx, updateTodo, arg.Title, arg.Description, arg.Completed, arg.ID).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdateAt,
	)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}
