package dto

import (
	"strconv"

	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/response"
)

type CreateCategoryRequest struct {
	Name     string `json:"name" binding:"required"`
	ParentId string `json:"parentId"`
	Color    string `json:"color" binding:"required"`
}

func (c CreateCategoryRequest) GetValidateMessage() response.ErrorMessages {
	return response.ErrorMessages{
		"Name.required":  "分类名称不能为空",
		"Color.required": "分类颜色不能为空",
	}
}

type CreateCategoryResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ParentId  string `json:"parentId"`
	Color     string `json:"color"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (r *CreateCategoryResponse) Transform(modal postgresql.Category) {
	r.ID = strconv.FormatInt(modal.ID, 10)
	r.Name = modal.Name
	r.ParentId = strconv.FormatInt(modal.ParentID, 10)
	r.Color = strconv.FormatInt(int64(modal.Color), 16)
	r.CreatedAt = modal.CreatedAt.Format("2006-01-02 15:04:05")
	r.UpdatedAt = modal.UpdatedAt.Format("2006-01-02 15:04:05")
}

type ModifyCategoryRequest struct {
	ID       string `uri:"id" binding:"required"`
	Name     string `json:"name"`
	ParentId string `json:"parentId"`
	Color    string `json:"color"`
}

func (r ModifyCategoryRequest) GetValidateMessage() response.ErrorMessages {
	return response.ErrorMessages{
		"ID.required": "ID不能为空",
	}
}

type FilterCategoryRequest struct {
	Name string `form:"name"`
}

func (c FilterCategoryRequest) GetValidateMessage() response.ErrorMessages {
	return response.ErrorMessages{}
}
