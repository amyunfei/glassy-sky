package handler

import (
	"github.com/amyunfei/glassy-sky/internal/admin/dto"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/response"
	"github.com/amyunfei/glassy-sky/internal/admin/service"
	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	Service service.UserService
}

func NewUserHandlers(service service.UserService) UserHandlers {
	return UserHandlers{Service: service}
}

// @Tags    用户信息
// @Summary 注册用户
// @Param   body    body     dto.CreateUserRequest  true "用户注册信息"
// @Success 200     {object} dto.CreateUserResponse
// @Router  /user/register  [POST]
func (h UserHandlers) RegisterUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	res, err := h.Service.CreateUser(ctx, req)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, res, "success")
}

// @Tags    用户信息
// @Summary 发送邮箱验证码
// @Param   email        path     string  true "用户注册信息"
// @Success 200          {object} dto.SuccessEmptyResponse
// @Router  /user/email  [GET]
func (h UserHandlers) SendEmailCode(ctx *gin.Context) {
	var req dto.SendEmailCodeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	err := h.Service.SendEmailCode(ctx, req)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, nil, "success")
}
