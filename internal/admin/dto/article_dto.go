package dto

import "github.com/amyunfei/glassy-sky/internal/admin/infrastructure/response"

type CreateArticleRequest struct {
	Title      string `json:"title" binding:"required"`
	Excerpt    string `json:"excerpt" binding:"required"`
	Content    string `json:"content" binding:"required"`
	CategoryId string `json:"categoryId" binding:"required"`
	LabelIds   string `json:"labelIds"`
}

func (c CreateArticleRequest) GetValidateMessage() response.ErrorMessages {
	return response.ErrorMessages{
		"Title.required":      "文章标题不能为空",
		"Excerpt.required":    "文章摘要不能为空",
		"Content.required":    "文章内容不能为空",
		"CategoryId.required": "文章分类不能为空",
	}
}

type CreateArticleResponse struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Excerpt   string   `json:"excerpt"`
	Content   string   `json:"content"`
	Category  string   `json:"category"`
	Labels    []string `json:"labels"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}
