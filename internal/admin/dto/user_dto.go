package dto

import (
	"strconv"

	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/response"
)

type SendEmailCodeRequest struct {
	Email string `uri:"email" binding:"required,email"`
}

func (c SendEmailCodeRequest) GetValidateMessage() response.ErrorMessages {
	return response.ErrorMessages{
		"Email.required": "邮箱地址不能为空",
		"Email.email":    "邮箱地址不正确",
	}
}

type CheckEmailCodeRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Code     string `json:"code" binding:"required"`
	Nickname string `json:"nickname"`
}

func (c CreateUserRequest) GetValidateMessage() response.ErrorMessages {
	return response.ErrorMessages{
		"Username.required": "用户名不能为空",
		"Password.required": "用户密码不能为空",
		"Email.required":    "邮箱地址不能为空",
		"Email.email":       "邮箱地址不正确",
		"Code.required":     "验证码不能为空",
	}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c LoginRequest) GetValidateMessage() response.ErrorMessages {
	return response.ErrorMessages{
		"Username.required": "用户名不能为空",
		"Password.required": "用户密码不能为空",
	}
}

type CreateUserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (r *CreateUserResponse) Transform(modal postgresql.User) {
	r.ID = strconv.FormatInt(modal.ID, 10)
	r.Username = modal.Username
	r.Nickname = modal.Nickname
	r.Email = modal.Email
	r.CreatedAt = modal.CreatedAt.Format("2006-01-02 15:04:05")
	r.UpdatedAt = modal.UpdatedAt.Format("2006-01-02 15:04:05")
}

type ModifyUserRequest struct {
	ID       string `uri:"id" binding:"required"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

func (r ModifyUserRequest) GetValidateMessage() response.ErrorMessages {
	return response.ErrorMessages{
		"ID.required": "ID不能为空",
	}
}

type FilterUserRequest struct {
	Username string `form:"username"`
	Nickname string `form:"nickname"`
}

func (c FilterUserRequest) GetValidateMessage() response.ErrorMessages {
	return response.ErrorMessages{}
}
