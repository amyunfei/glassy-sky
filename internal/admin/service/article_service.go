package service

import (
	"context"
	"time"

	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/dto"
)

type ArticleService interface {
	CreateArticle(ctx context.Context, data dto.CreateArticleRequest) (*dto.CreateArticleResponse, error)
}
type ArticleOptions interface {
	GetTimeZone() *time.Location
}
type DefaultArticleService struct {
	repo    postgresql.Repository
	options ArticleOptions
}

func NewArticleService(repo postgresql.Repository, options ArticleOptions) ArticleService {
	return &DefaultArticleService{repo, options}
}

func (s DefaultArticleService) CreateArticle(ctx context.Context, data dto.CreateArticleRequest) (*dto.CreateArticleResponse, error) {
	return nil, nil
}
