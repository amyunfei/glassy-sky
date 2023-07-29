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
// @Param   body    body     dto.RegisterUserRequest  true "用户注册信息"
// @Success 200     {object} response.Body[dto.CreateUserResponse]
// @Router  /user/register  [POST]
func (h UserHandlers) RegisterUser(ctx *gin.Context) {
	var req dto.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	validity := h.Service.CheckEmailCode(ctx, dto.CheckEmailCodeRequest{
		Email: req.Email,
		Code:  req.Code,
	})
	if !validity {
		response.RequestError(ctx, "邮箱验证失败")
		return
	}
	res, err := h.Service.CreateUser(ctx, dto.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Nickname: req.Email,
	})
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, res, "success")
}

// @Tags    用户信息
// @Summary 验证邮箱可用
// @Param   email        path          string  true "邮箱地址"
// @Success 200          {object}      dto.SuccessEmptyResponse
// @Router  /user/email-verify/{email} [GET]
func (h UserHandlers) VerifyEmail(ctx *gin.Context) {
	var req dto.SendEmailCodeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	avaliable, err := h.Service.VerifyEmail(ctx, req)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, avaliable, "success")
}

// @Tags    用户信息
// @Summary 发送邮箱验证码
// @Param   email        path     string  true "邮箱地址"
// @Success 200          {object} dto.SuccessEmptyResponse
// @Router  /user/email-code/{email}  [GET]
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
	response.Success(ctx, "", "success")
}

// @Tags    用户信息
// @Summary 创建用户
// @Param   data     body     dto.CreateUserRequest                true "用户信息"
// @Success 200     {object} response.Body[dto.CreateUserResponse]
// @Router  /user   [POST]
func (h UserHandlers) CreateUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	user, err := h.Service.CreateUser(ctx, req)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, user, "success")
}

// @Tags    用户信息
// @Summary 登录
// @Param   data        body     dto.LoginRequest       true "登录信息"
// @Success 200         {object} response.Body[string]
// @Router  /user/login [POST]
func (h UserHandlers) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	token, err := h.Service.Login(ctx, req)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, token, "success")
}

// @Tags    用户信息
// @Summary 查询用户
// @Param   id          path     string                   true "用户id"
// @Success 200         {object} response.Body[dto.CreateUserResponse]
// @Router  /user/{id}  [GET]
func (h UserHandlers) GetUser(ctx *gin.Context) {
	var req dto.UriIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	user, err := h.Service.GetUser(ctx, req)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, user, "success")
}

// @Tags    用户信息
// @Summary 修改用户信息
// @Param   data       body     dto.ModifyUserRequest      true "用户信息"
// @Success 200        {object} response.Body[dto.CreateUserResponse]
// @Router  /user/{id} [PUT]
func (h UserHandlers) ModifyUser(ctx *gin.Context) {
	var req dto.ModifyUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	result, err := h.Service.ModifyUser(ctx, req)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, result, "success")
}

// @Tags    用户信息
// @Summary 删除用户
// @Param   id          path     string                   true "用户id"
// @Success 200         {object} dto.SuccessEmptyResponse
// @Router  /user/{id}  [DELETE]
func (h UserHandlers) DeleteUser(ctx *gin.Context) {
	var req dto.UriIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ValidationError(ctx, req, err)
		return
	}
	err := h.Service.DeleteUser(ctx, req)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, "", "success")
}

// @Tags    用户信息
// @Summary 分页查询用户
// @Param   pageParams query    dto.ListRequest            true "分页参数"
// @Param   filter     query    dto.FilterUserRequest      true "筛选参数"
// @Success 200        {object} response.Body[dto.ListResponse[dto.CreateUserResponse]]
// @Router  /user      [GET]
func (h UserHandlers) ListUser(ctx *gin.Context) {
	var listReq dto.ListRequest
	if err := ctx.ShouldBindQuery(&listReq); err != nil {
		response.ValidationError(ctx, listReq, err)
		return
	}
	var filterReq dto.FilterUserRequest
	if err := ctx.ShouldBindQuery(&filterReq); err != nil {
		response.ValidationError(ctx, filterReq, err)
		return
	}
	result, err := h.Service.ListUser(ctx, listReq, filterReq)
	if err != nil {
		response.UnexpectedError(ctx, err.Error())
		return
	}
	response.Success(ctx, result, "success")
}
