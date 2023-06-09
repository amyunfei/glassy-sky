package postgresql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomArticleLabel(t *testing.T) ArticlesLabel {
	label := createRandomLabel(t)
	article := createRandomArticle(t)
	arg := CreateArticleLabelParams{
		ArticleID: article.ID,
		LabelID:   label.ID,
	}
	articlesLabel, err := testQueries.CreateArticleLabel(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, articlesLabel)

	require.NotZero(t, articlesLabel.LabelID)
	require.NotZero(t, articlesLabel.ArticleID)

	require.Equal(t, arg.LabelID, articlesLabel.LabelID)
	require.Equal(t, arg.ArticleID, articlesLabel.ArticleID)
	return articlesLabel
}

func TestCreateArticleLabel(t *testing.T) {
	createRandomArticleLabel(t)
}

func TestGetArticleLabel(t *testing.T) {
	articlesLabel1 := createRandomArticleLabel(t)
	arg := GetArticleLabelParams(articlesLabel1)
	articlesLabel2, err := testQueries.GetArticleLabel(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, articlesLabel2)

	require.Equal(t, articlesLabel1.LabelID, articlesLabel2.LabelID)
	require.Equal(t, articlesLabel1.ArticleID, articlesLabel2.ArticleID)
}

func TestGetArticlesByLabelID(t *testing.T) {
	articlesLabel1 := createRandomArticleLabel(t)
	articlesLabelGroup, err := testQueries.GetArticlesByLabelID(context.Background(), articlesLabel1.LabelID)
	require.NoError(t, err)
	require.NotEmpty(t, articlesLabelGroup)
	for _, ac := range articlesLabelGroup {
		require.Equal(t, articlesLabel1.LabelID, ac.LabelID)
	}
}

func TestGetLabelsByArticleID(t *testing.T) {
	articlesLabel1 := createRandomArticleLabel(t)
	articlesLabelGroup, err := testQueries.GetLabelsByArticleID(context.Background(), articlesLabel1.ArticleID)
	require.NoError(t, err)
	require.NotEmpty(t, articlesLabelGroup)
	for _, ac := range articlesLabelGroup {
		require.Equal(t, articlesLabel1.ArticleID, ac.ArticleID)
	}
}

func TestDeleteArticleLabel(t *testing.T) {
	articlesLabel1 := createRandomArticleLabel(t)
	arg1 := DeleteArticleLabelParams(articlesLabel1)
	err := testQueries.DeleteArticleLabel(context.Background(), arg1)
	require.NoError(t, err)
	arg2 := GetArticleLabelParams(articlesLabel1)
	articlesLabel2, err := testQueries.GetArticleLabel(context.Background(), arg2)
	require.Error(t, err)
	require.Empty(t, articlesLabel2)
}
