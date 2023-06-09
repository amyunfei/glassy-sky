package postgresql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomArticleCategory(t *testing.T) ArticlesCategory {
	category := createRandomCategory(t, DEFAULT_PARENT_ID)
	article := createRandomArticle(t)
	arg := CreateArticleCategoryParams{
		ArticleID:  article.ID,
		CategoryID: category.ID,
	}
	articleCategory, err := testQueries.CreateArticleCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, articleCategory)

	require.NotZero(t, articleCategory.CategoryID)
	require.NotZero(t, articleCategory.ArticleID)

	require.Equal(t, arg.CategoryID, articleCategory.CategoryID)
	require.Equal(t, arg.ArticleID, articleCategory.ArticleID)
	return articleCategory
}
func TestCreateArticleCategory(t *testing.T) {
	createRandomArticleCategory(t)
}

func TestGetArticlesByCategoryID(t *testing.T) {
	articlesCategory1 := createRandomArticleCategory(t)
	articlesCategoryGroup, err := testQueries.GetArticlesByCategoryID(context.Background(), articlesCategory1.CategoryID)
	require.NoError(t, err)
	require.NotEmpty(t, articlesCategoryGroup)
	for _, ac := range articlesCategoryGroup {
		require.Equal(t, articlesCategory1.CategoryID, ac.CategoryID)
	}
}

func TestGetCategoriesByArticleID(t *testing.T) {
	articlesCategory1 := createRandomArticleCategory(t)
	articlesCategoryGroup, err := testQueries.GetCategoriesByArticleID(context.Background(), articlesCategory1.ArticleID)
	require.NoError(t, err)
	require.NotEmpty(t, articlesCategoryGroup)
	for _, ac := range articlesCategoryGroup {
		require.Equal(t, articlesCategory1.ArticleID, ac.ArticleID)
	}
}

func TestDeleteArticleCategory(t *testing.T) {
	articleCategory1 := createRandomArticleCategory(t)
	arg1 := DeleteArticleCategoryParams(articleCategory1)
	err := testQueries.DeleteArticleCategory(context.Background(), arg1)
	require.NoError(t, err)
	arg2 := GetArticleCategoryParams(articleCategory1)
	articleCategory2, err := testQueries.GetArticleCategory(context.Background(), arg2)
	require.Error(t, err)
	require.Empty(t, articleCategory2)
}
