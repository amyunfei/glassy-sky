package dto

import (
	"strconv"

	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/response"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/utils"
)

type CreateLabelRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

func (c CreateLabelRequest) GetValidateMessage() response.ErrorMessages {
	return response.ErrorMessages{
		"Name.required":  "标签名称不能为空",
		"Color.required": "标签颜色不能为空",
	}
}

type CreateLabelResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (r *CreateLabelResponse) Transform(modal postgresql.Label) {
	r.ID = strconv.FormatInt(modal.ID, 10)
	r.Name = modal.Name
	r.Color = utils.IntToHexColor(modal.Color)
	r.CreatedAt = modal.CreatedAt.Format("2006-01-02 15:04:05")
	r.UpdatedAt = modal.UpdatedAt.Format("2006-01-02 15:04:05")
}

type ModifyLabelRequest struct {
	ID    string `uri:"id" binding:"required"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (r ModifyLabelRequest) GetValidateMessage() response.ErrorMessages {
	return response.ErrorMessages{
		"ID.required": "ID不能为空",
	}
}

type FilterLabelRequest struct {
	Name string `form:"name"`
}

func (c FilterLabelRequest) GetValidateMessage() response.ErrorMessages {
	return response.ErrorMessages{}
}
