package mocks

import (
	"context"
	"github.com/abdukhashimov/go_api_mono_repo/internal/core/repository/psql/sqlc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func createToDo(t *testing.T) sqlc.Todo {
	ctx := context.Background()
	arg := sqlc.CreateTodoParams{
		Title:       "test",
		Description: "test description",
	}
	todo, err := testQueries.CreateTodo(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	return todo
}

func TestCreateToDo(t *testing.T) {
	createToDo(t)
}

func TestGetToDoById(t *testing.T) {
	todo := createToDo(t)
	fetchedTodo, err := testQueries.GetTodoById(context.Background(), todo.ID)
	require.NoError(t, err)
	require.Equal(t, todo.Title, fetchedTodo.Title)
	require.Equal(t, todo.Description, fetchedTodo.Description)
	require.Equal(t, todo.Completed, fetchedTodo.Completed)
}

func updateTodo(t *testing.T, id int64) sqlc.Todo {
	ctx := context.Background()
	arg := sqlc.UpdateTodoParams{
		ID:          id,
		Title:       "test",
		Description: "test description",
		Completed:   true,
	}
	todo, err := testQueries.UpdateTodo(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo)
	assert.Equal(t, arg.Title, todo.Title)
	assert.Equal(t, arg.Description, todo.Description)
	assert.Equal(t, arg.Completed, todo.Completed)
	assert.NotEqual(t, arg.Completed, false)

	return todo
}

func TestUpdateTodoById(t *testing.T) {
	todo := createToDo(t)
	updateTodo(t, todo.ID)
}
