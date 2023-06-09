package postgresql

import (
	"context"
	"testing"
	"time"

	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/utils"
	"github.com/stretchr/testify/require"
)

func createRandomArticle(t *testing.T) Article {
	user := createRandomUser(t)
	arg := CreateArticleParams{
		Title:   utils.RandomString(6),
		Excerpt: utils.RandomString(100),
		Content: utils.RandomString(1000),
		UserID:  user.ID,
	}
	currentTime := time.Now()
	article, err := testQueries.CreateArticle(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, article)

	require.Equal(t, arg.Title, article.Title)
	require.Equal(t, arg.Excerpt, article.Excerpt)
	require.Equal(t, arg.Content, article.Content)
	require.Equal(t, arg.UserID, article.UserID)

	require.NotZero(t, article.ID)
	require.NotZero(t, article.CreatedAt)
	require.NotZero(t, article.UpdatedAt)
	require.Empty(t, article.DeletedAt)

	require.WithinDuration(t, currentTime, article.CreatedAt, time.Second)
	require.WithinDuration(t, currentTime, article.UpdatedAt, time.Second)
	return article
}

func TestCreateArticle(t *testing.T) {
	createRandomArticle(t)
}

func TestGetArticle(t *testing.T) {
	article1 := createRandomArticle(t)
	article2, err := testQueries.GetArticle(context.Background(), article1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, article2)

	require.Equal(t, article1.ID, article2.ID)
	require.Equal(t, article1.Title, article2.Title)
	require.Equal(t, article1.Excerpt, article2.Excerpt)
	require.Equal(t, article1.Content, article2.Content)
	require.Equal(t, article1.UserID, article2.UserID)
	require.Equal(t, article1.CreatedAt, article2.CreatedAt)
	require.Equal(t, article1.UpdatedAt, article2.UpdatedAt)
}

func TestDeleteArticle(t *testing.T) {
	article1 := createRandomArticle(t)
	err := testQueries.DeleteArticle(context.Background(), article1.ID)
	require.NoError(t, err)
	article2, err := testQueries.GetArticle(context.Background(), article1.ID)
	require.Error(t, err)
	require.Empty(t, article2)
}

func TestCountArticle(t *testing.T) {
	count1, err := testQueries.CountArticle(context.Background(), "")
	require.NoError(t, err)
	createRandomArticle(t)
	count2, err := testQueries.CountArticle(context.Background(), "")
	require.NoError(t, err)
	require.Equal(t, count1+1, count2)
}

func TestListArticle(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomArticle(t)
	}
	arg := ListArticleParams{
		Limit:  5,
		Offset: 5,
		Title:  "",
	}
	articles, err := testQueries.ListArticle(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, articles, 5)
	for _, article := range articles {
		require.NotEmpty(t, article)
	}
}
