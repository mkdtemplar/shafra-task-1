package db

import (
	"context"
	"shafra-task-1/internal/models"
	"shafra-task-1/utils"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func createRandomUSer(t *testing.T) *models.User {
	testRepo := NewUserRepo(conn)
	arg := &models.User{
		ID:          utils.RandomInt(1, 100),
		NameSurname: utils.RandomString(),
		Age:         utils.RandomInt(20, 100),
	}

	user, err := testRepo.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, arg, user)

	return user
}

func TestPostgresDB_CreateUser(t *testing.T) {
	createRandomUSer(t)
}

func TestPostgresDB_GetUserById(t *testing.T) {
	testRepo := NewUserRepo(conn)
	user1 := createRandomUSer(t)
	user2, err := testRepo.GetUserById(context.Background(), user1.ID)
	require.NoError(t, err)
	require.Equal(t, user1.NameSurname, user2.NameSurname)
	require.Equal(t, user1.Age, user2.Age)

}

func TestPostgresDB_UpdateUser(t *testing.T) {
	testRepo := NewUserRepo(conn)
	user1 := createRandomUSer(t)

	args := &models.User{
		ID:          user1.ID,
		NameSurname: utils.RandomString(),
		Age:         utils.RandomInt(20, 100),
	}

	updatedUser, err := testRepo.UpdateUser(context.Background(), user1.ID, args.NameSurname, int(args.Age))
	require.NoError(t, err)
	require.Equal(t, updatedUser.NameSurname, args.NameSurname)
	require.Equal(t, updatedUser.Age, args.Age)

}

func TestPostgresDB_DeleteUser(t *testing.T) {
	testRepo := NewUserRepo(conn)
	user1 := createRandomUSer(t)
	err := testRepo.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)
	user2, err := testRepo.GetUserById(context.Background(), user1.ID)
	require.Error(t, gorm.ErrRecordNotFound, err)
	require.Empty(t, user2)

}
