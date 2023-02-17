package handler

import (
	"github.com/amyunfei/glassy-sky/internal/admin/dto"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/response"
	"github.com/amyunfei/glassy-sky/internal/admin/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandlers struct {
	Service service.CategoryService
}

// @Tags    分类信息
// @Summary 创建分类
// @Param   body      body     dto.CreateCategoryRequest  true "分类信息"
// @Success 200       {object} dto.CreateCategoryResponse
// @Router  /category [POST]
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

// @Tags    分类信息
// @Summary 删除分类
// @Param   id             path     string                   true "分类id"
// @Success 200            {object} dto.SuccessEmptyResponse
// @Router  /category/{id} [DELETE]
func (h CategoryHandlers) DeleteCategory(ctx *gin.Context) {
	var req dto.UriIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	err := h.Service.DeleteCategory(ctx, req)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, nil, "success")
}

// @Tags    分类信息
// @Summary 修改分类
// @Param   body           body     dto.ModifyCategoryRequest  true "分类信息"
// @Success 200            {object} dto.CreateCategoryResponse
// @Router  /category/{id} [PUT]
func (h CategoryHandlers) ModifyCategory(ctx *gin.Context) {
	var req dto.ModifyCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	res, err := h.Service.ModifyCategory(ctx, req)
	if err != nil {
		logger.Error(err.Error())
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, res, "success")
}

// @Tags    分类信息
// @Summary 分页查询分类
// @Param   pageParams query    dto.ListRequest           true "分页参数"
// @Param   filter     query    dto.FilterCategoryRequest true "筛选参数"
// @Success 200        {object} dto.CreateCategoryResponse
// @Router  /category  [GET]
func (h CategoryHandlers) ListCategory(ctx *gin.Context) {
	var listReq dto.ListRequest
	if err := ctx.ShouldBindQuery(&listReq); err != nil {
		response.ValidationError(ctx, listReq, err)
		return
	}
	var filterReq dto.FilterCategoryRequest
	if err := ctx.ShouldBindQuery(&filterReq); err != nil {
		response.ValidationError(ctx, filterReq, err)
		return
	}
	res, err := h.Service.ListCategory(ctx, listReq, filterReq)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, res, "success")
}
