package mocks

import (
	"context"
	"github.com/abdukhashimov/go_api_mono_repo/internal/core/repository/psql/sqlc"
	"github.com/stretchr/testify/require"
	"testing"
)

func createUser(t *testing.T) sqlc.User {
	ctx := context.Background()
	arg := sqlc.CreateUserParams{
		Name:  "test",
		Email: "test@mail.com",
	}

	err := testQueries.CreateUser(ctx, arg)
	require.NoError(t, err)

	user, err := testQueries.GetUserByName(ctx, arg.Name)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}

func TestCreateUser(t *testing.T) {
	createUser(t)
}

func TestGetUserById(t *testing.T) {
	user := createUser(t)
	fetchedUser, err := testQueries.GetUserById(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, user.Name, fetchedUser.Name)
}
