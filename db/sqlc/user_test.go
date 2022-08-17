package db

import (
	"context"
	"simplebank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	// create user first
	randomUser := createRandomUser(t)
	result, err := testQueries.GetUser(context.Background(), randomUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, randomUser.Username, result.Username)
	require.Equal(t, randomUser.HashedPassword, result.HashedPassword)
	require.Equal(t, randomUser.FullName, result.FullName)
	require.Equal(t, randomUser.Email, result.Email)
	require.WithinDuration(t, randomUser.PasswordChangedAt, result.PasswordChangedAt, time.Second)
	require.WithinDuration(t, randomUser.CreatedAt, result.CreatedAt, time.Second)
}
