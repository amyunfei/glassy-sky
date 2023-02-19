package postgresql

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/utils"
	"github.com/stretchr/testify/require"
)

func createRandomLabel(t *testing.T) Label {
	arg := CreateLabelParams{
		Name:  utils.RandomString(5),
		Color: utils.RandomColorInt(),
	}

	currentTime := time.Now()
	label, err := testQueries.CreateLabel(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, label)

	require.Equal(t, arg.Name, label.Name)
	require.Equal(t, arg.Color, label.Color)

	require.NotZero(t, label.ID)
	require.NotZero(t, label.CreatedAt)
	require.NotZero(t, label.UpdatedAt)
	require.Empty(t, label.DeletedAt)

	require.Equal(t, label.CreatedAt, label.UpdatedAt)
	require.WithinDuration(t, currentTime, label.CreatedAt, time.Second)
	require.WithinDuration(t, currentTime, label.UpdatedAt, time.Second)
	return label
}

func TestCreateLabel(t *testing.T) {
	createRandomLabel(t)
}

func TestGetLabel(t *testing.T) {
	label1 := createRandomLabel(t)
	label2, err := testQueries.GetLabel(context.Background(), label1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, label2)

	require.Equal(t, label1.ID, label2.ID)
	require.Equal(t, label1.Name, label2.Name)
	require.Equal(t, label1.Color, label2.Color)
	require.Equal(t, label1.CreatedAt, label2.CreatedAt)
	require.Equal(t, label1.UpdatedAt, label2.UpdatedAt)
}

func TestUpdateLabel(t *testing.T) {
	label1 := createRandomLabel(t)
	arg := UpdateLabelParams{
		ID: label1.ID,
		Name: sql.NullString{
			String: utils.RandomString(5),
			Valid:  true,
		},
		Color: sql.NullInt32{
			Int32: utils.RandomColorInt(),
			Valid: true,
		},
	}
	label2, err := testQueries.UpdateLabel(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, label2)

	require.Equal(t, label1.ID, label2.ID)
	require.Equal(t, label1.CreatedAt, label2.CreatedAt)

	require.Equal(t, arg.Name.String, label2.Name)
	require.Equal(t, arg.Color.Int32, label2.Color)

	require.WithinDuration(t, label1.UpdatedAt, label2.UpdatedAt, time.Second)
}

func TestCountLabel(t *testing.T) {
	count1, err := testQueries.CountLabel(context.Background(), "")
	require.NoError(t, err)
	createRandomLabel(t)
	count2, err := testQueries.CountLabel(context.Background(), "")
	require.NoError(t, err)
	require.Equal(t, count1+1, count2)
}

func TestListLabel(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomLabel(t)
	}
	arg := ListLabelParams{
		Limit:  5,
		Offset: 5,
		Name:   "",
	}
	labels, err := testQueries.ListLabel(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, labels, 5)
	for _, label := range labels {
		require.NotEmpty(t, label)
	}
}
