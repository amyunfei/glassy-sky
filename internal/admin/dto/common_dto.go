package dto

import "github.com/amyunfei/glassy-sky/internal/admin/infrastructure/response"

type UriIdRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (c UriIdRequest) GetValidateMessage() response.ErrorMessages {
	return response.ErrorMessages{
		"ID.required": "ID不能为空",
	}
}

type ListRequest struct {
	Size    int32 `form:"size" binding:"required"`
	Current int32 `form:"current" binding:"required"`
}

func (r ListRequest) GetValidateMessage() response.ErrorMessages {
	return response.ErrorMessages{
		"Size.required":    "每页条数不能为空",
		"Current.required": "当前页码不能为空",
	}
}

type ListResponse[T any] struct {
	List  []T   `json:"list"`
	Count int64 `json:"count"`
}

type SuccessEmptyResponse struct {
	Code    int    `json:"code" example:"0"`
	Data    string `json:"data" example:""`
	Message string `json:"msg" example:"success"`
}
