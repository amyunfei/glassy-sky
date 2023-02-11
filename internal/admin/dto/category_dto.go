package dto

import (
	"reflect"
	"strconv"

	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/go-playground/validator"
)

type CreateCategoryRequest struct {
	Name     string `json:"name" binding:"required"`
	ParentId string `json:"parentId"`
	Color    string `json:"color" binding:"required"`
}

type CreateCategoryResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ParentId  string `json:"parentId"`
	Color     string `json:"color"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (c *CreateCategoryResponse) Transform(modal postgresql.Category) {
	c.ID = strconv.FormatInt(modal.ID, 10)
	c.Name = modal.Name
	c.ParentId = strconv.FormatInt(modal.ParentID, 10)
	c.Color = strconv.FormatInt(int64(modal.Color), 16)
	c.CreatedAt = modal.CreatedAt.Format("2006-01-02 15:04:05")
	c.UpdatedAt = modal.UpdatedAt.Format("2006-01-02 15:04:05")
}

// 自定义错误消息
func GetError(errs validator.ValidationErrors, r interface{}) string {
	s := reflect.TypeOf(r)
	for _, fieldError := range errs {
		filed, _ := s.FieldByName(fieldError.Field())
		errTag := fieldError.Tag() + "_err"
		// 获取对应binding得错误消息
		errTagText := filed.Tag.Get(errTag)
		// 获取统一错误消息
		errText := filed.Tag.Get("err")
		if errTagText != "" {
			return errTagText
		}
		if errText != "" {
			return errText
		}
		return fieldError.Field() + ":" + fieldError.Tag()
	}
	return ""
}
