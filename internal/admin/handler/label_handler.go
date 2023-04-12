package handler

import (
	"github.com/amyunfei/glassy-sky/internal/admin/dto"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/response"
	"github.com/amyunfei/glassy-sky/internal/admin/service"
	"github.com/gin-gonic/gin"
)

type LabelHandlers struct {
	Service service.LabelService
}

func NewLabelHandlers(service service.LabelService) LabelHandlers {
	return LabelHandlers{Service: service}
}

// @Tags    标签信息
// @Summary 创建标签
// @Param   body    body     dto.CreateLabelRequest  true "标签信息"
// @Success 200     {object} response.Body[dto.CreateLabelResponse]
// @Router  /label  [POST]
func (h LabelHandlers) CreateLabel(ctx *gin.Context) {
	var req dto.CreateLabelRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	res, err := h.Service.CreateLabel(ctx, req)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, res, "success")
}

// @Tags    标签信息
// @Summary 删除标签
// @Param   id          path     string                   true "标签id"
// @Success 200         {object} dto.SuccessEmptyResponse
// @Router  /label/{id} [DELETE]
func (h LabelHandlers) DeleteLabel(ctx *gin.Context) {
	var req dto.UriIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	err := h.Service.DeleteLabel(ctx, req)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, "", "success")
}

// @Tags    标签信息
// @Summary 修改标签
// @Param   body        body     dto.ModifyLabelRequest  true "标签信息"
// @Success 200         {object} response.Body[dto.CreateLabelResponse]
// @Router  /label/{id} [PUT]
func (h LabelHandlers) ModifyLabel(ctx *gin.Context) {
	var req dto.ModifyLabelRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	res, err := h.Service.ModifyLabel(ctx, req)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, res, "success")
}

// @Tags    标签信息
// @Summary 分页查询标签
// @Param   pageParams query    dto.ListRequest            true "分页参数"
// @Param   filter     query    dto.FilterLabelRequest     true "筛选参数"
// @Success 200        {object} response.Body[dto.ListResponse[dto.CreateLabelResponse]]
// @Router  /label     [GET]
func (h LabelHandlers) ListLabel(ctx *gin.Context) {
	var listReq dto.ListRequest
	if err := ctx.ShouldBindQuery(&listReq); err != nil {
		response.ValidationError(ctx, listReq, err)
		return
	}
	var filterReq dto.FilterLabelRequest
	if err := ctx.ShouldBindQuery(&filterReq); err != nil {
		response.ValidationError(ctx, filterReq, err)
		return
	}
	res, err := h.Service.ListLabel(ctx, listReq, filterReq)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, res, "success")
}

// @Tags    标签信息
// @Summary 获取标签
// @Param   id          path     string                   true "标签id"
// @Success 200         {object} response.Body[dto.CreateLabelResponse]
// @Router  /label/{id} [GET]
func (h LabelHandlers) GetLabel(ctx *gin.Context) {
	var req dto.UriIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	res, err := h.Service.GetLabel(ctx, req)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, res, "success")
}
