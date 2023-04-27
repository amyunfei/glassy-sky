package postgresql

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/utils"
	"github.com/stretchr/testify/require"
)

const DEFAULT_PARENT_ID = 0

func createRandomCategory(t *testing.T, parentId int64) Category {
	arg := CreateCategoryParams{
		Name:     utils.RandomString(6),
		ParentID: parentId,
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
	createRandomCategory(t, DEFAULT_PARENT_ID)
}

func TestGetCategory(t *testing.T) {
	category1 := createRandomCategory(t, DEFAULT_PARENT_ID)
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
	category1 := createRandomCategory(t, DEFAULT_PARENT_ID)
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

func TestCountCategory(t *testing.T) {
	count1, err := testQueries.CountCategory(context.Background(), "")
	require.NoError(t, err)
	createRandomCategory(t, DEFAULT_PARENT_ID)
	count2, err := testQueries.CountCategory(context.Background(), "")
	require.NoError(t, err)
	require.Equal(t, count1+1, count2)
}

func TestListCategory(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCategory(t, DEFAULT_PARENT_ID)
	}
	arg := ListCategoryParams{
		Limit:  5,
		Offset: 5,
		Name:   "",
	}
	categories, err := testQueries.ListCategory(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, categories, 5)
	for _, category := range categories {
		require.NotEmpty(t, category)
	}
}

func TestTreeCategory(t *testing.T) {
	var parentId int64 = DEFAULT_PARENT_ID
	for i := 0; i < 10; i++ {
		category := createRandomCategory(t, parentId)
		parentId = category.ID
	}
	categories, err := testQueries.TreeCategory(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, categories)
	for _, category := range categories {
		require.NotEmpty(t, category)
		require.NotEmpty(t, category.Level)
		if category.Level > 1 {
			require.NotEmpty(t, category.ParentID)
			parentCategory, err := testQueries.GetCategory(context.Background(), category.ParentID)
			require.NoError(t, err)
			require.NotEmpty(t, parentCategory)
			require.Equal(t, parentCategory.ID, category.ParentID)
		}
	}
}
