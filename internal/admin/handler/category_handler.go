package handler

import (
	"github.com/amyunfei/glassy-sky/internal/admin/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandlers struct {
	Service service.CategoryService
}

func (h CategoryHandlers) Create(ctx *gin.Context) {
	// h.service.CreateCategory()
}
