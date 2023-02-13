package handler

import (
	"github.com/amyunfei/glassy-sky/internal/admin/dto"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/response"
	"github.com/amyunfei/glassy-sky/internal/admin/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandlers struct {
	Service service.CategoryService
}

// @Tags 分类信息
// @Summary 创建分类
// @Param body body dto.CreateCategoryRequest true "分类信息"
// @Success 200 {object} dto.CreateCategoryResponse
// @Router /category [POST]
func (h CategoryHandlers) CreateCategory(ctx *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	res, err := h.Service.CreateCategory(ctx, req)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, res, "success")
}
