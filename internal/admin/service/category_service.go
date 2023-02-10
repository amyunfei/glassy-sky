package service

import "github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"

type CategoryService interface {
	CreateCategory()
}

type DefaultCategoryService struct {
	repo *postgresql.Queries
}

func (s DefaultCategoryService) CreateCategory() {

}

func NewCategoryService(repo *postgresql.Queries) DefaultCategoryService {
	return DefaultCategoryService{repo}
}
