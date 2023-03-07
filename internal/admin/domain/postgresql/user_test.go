package postgresql

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: utils.RandomString(5),
		Password: utils.RandomString(20),
		Email:    utils.RandomEmail(),
		Nickname: utils.RandomString(8),
	}

	currentTime := time.Now()
	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Nickname, user.Nickname)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
	require.Empty(t, user.DeletedAt)

	require.WithinDuration(t, currentTime, user.CreatedAt, time.Second)
	require.WithinDuration(t, currentTime, user.UpdatedAt, time.Second)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID: user1.ID,
		Nickname: sql.NullString{
			String: utils.RandomString(8),
			Valid:  true,
		},
		Avatar: sql.NullString{
			String: "",
			Valid:  false,
		},
		Password: sql.NullString{
			String: utils.RandomString(20),
			Valid:  true,
		},
	}
	currentTime := time.Now()
	user2, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)

	require.Equal(t, user1.Avatar, user2.Avatar)
	require.Equal(t, arg.Nickname.String, user2.Nickname)
	require.Equal(t, arg.Password.String, user2.Password)

	require.NotEqual(t, user1.Nickname, user2.Nickname)
	require.NotEqual(t, user1.Password, user2.Password)
	require.NotEqual(t, user1.UpdatedAt, user2.UpdatedAt)
	require.WithinDuration(t, currentTime, user2.UpdatedAt, time.Second)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Nickname, user2.Nickname)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.Equal(t, user1.UpdatedAt, user2.UpdatedAt)
}

func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByEmail(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Nickname, user2.Nickname)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.Equal(t, user1.UpdatedAt, user2.UpdatedAt)
}

func TestCountUser(t *testing.T) {
	arg := CountUserParams{
		Username: "",
		Nickname: "",
	}
	count1, err := testQueries.CountUser(context.Background(), arg)
	require.NoError(t, err)
	createRandomUser(t)
	count2, err := testQueries.CountUser(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, count1+1, count2)
}

func TestListUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}
	arg := ListUserParams{
		Limit:    5,
		Offset:   5,
		Username: "",
		Nickname: "",
	}
	users, err := testQueries.ListUser(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)
	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
