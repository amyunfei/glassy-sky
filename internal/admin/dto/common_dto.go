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

type SuccessEmptyResponse struct {
	Code    int    `json:"code" example:"0"`
	Data    string `json:"data" example:""`
	Message string `json:"msg" example:"success"`
}
