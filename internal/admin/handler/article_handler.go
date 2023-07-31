package handler

import (
	"github.com/amyunfei/glassy-sky/internal/admin/dto"
	"github.com/amyunfei/glassy-sky/internal/admin/service"
	"github.com/gin-gonic/gin"
)

type ArticleHandlers struct {
	Service service.ArticleService
}

func NewArticleHandlers(service service.ArticleService) ArticleHandlers {
	return ArticleHandlers{Service: service}
}

func (h ArticleHandlers) CreateArticle(ctx *gin.Context) {
	var req dto.CreateArticleRequest
	h.Service.CreateArticle(ctx, req)
}
