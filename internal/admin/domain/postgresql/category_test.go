package postgresql

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/utils"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	arg := CreateCategoryParams{
		Name:     utils.RandomString(6),
		ParentID: 0,
		Color:    utils.RandomColorInt(),
	}
	currentTime := time.Now()
	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Name, category.Name)
	require.Equal(t, arg.ParentID, category.ParentID)
	require.Equal(t, arg.Color, category.Color)

	require.NotZero(t, category.ID)
	require.NotZero(t, category.CreatedAt)
	require.NotZero(t, category.UpdatedAt)
	require.Empty(t, category.DeletedAt)

	require.WithinDuration(t, currentTime, category.CreatedAt, time.Second)
	require.WithinDuration(t, currentTime, category.UpdatedAt, time.Second)
	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
	require.Equal(t, category1.ParentID, category2.ParentID)
	require.Equal(t, category1.Color, category2.Color)
	require.Equal(t, category1.CreatedAt, category2.CreatedAt)
	require.Equal(t, category1.UpdatedAt, category2.UpdatedAt)
}

func TestUpdateCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	arg := UpdateCategoryParams{
		ID: category1.ID,
		Name: sql.NullString{
			String: utils.RandomString(6),
			Valid:  true,
		},
		Color: sql.NullInt32{
			Int32: utils.RandomColorInt(),
			Valid: true,
		},
	}
	category2, err := testQueries.UpdateCategory(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.ParentID, category2.ParentID)
	require.Equal(t, category1.CreatedAt, category2.CreatedAt)

	require.Equal(t, arg.Name.String, category2.Name)
	require.Equal(t, arg.Color.Int32, category2.Color)

	require.WithinDuration(t, category1.UpdatedAt, category2.UpdatedAt, time.Second)
}
