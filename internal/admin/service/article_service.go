package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/dto"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/token"
	"github.com/amyunfei/glassy-sky/internal/admin/middleware"
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
	var err error
	tokenPayload := ctx.Value(middleware.AuthorizationPayloadKey)
	payload, ok := tokenPayload.(*token.Payload)
	if !ok {
		err = errors.New("invalid token payload")
		logger.Error(err.Error())
		return nil, err
	}
	fmt.Println(payload.UserId)
	return nil, nil
}
