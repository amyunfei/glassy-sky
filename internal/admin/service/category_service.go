package service

import (
	"context"
	"strconv"

	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/dto"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, data dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, error)
}

type DefaultCategoryService struct {
	repo *postgresql.Queries
}

func (s DefaultCategoryService) CreateCategory(
	ctx context.Context,
	data dto.CreateCategoryRequest,
) (*dto.CreateCategoryResponse, error) {
	var parentId int64
	var err error
	if data.ParentId == "" {
		parentId = 0
	} else if parentId, err = strconv.ParseInt(data.ParentId, 10, 64); err != nil {
		logger.Error(err.Error())
		return nil, err
	} else {
		_, err = s.repo.GetCategory(ctx, parentId)
	}
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	color, err := strconv.ParseInt(data.Color, 16, 32)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	arg := postgresql.CreateCategoryParams{
		Name:     data.Name,
		ParentID: parentId,
		Color:    int32(color),
	}
	category, err := s.repo.CreateCategory(ctx, arg)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	var result dto.CreateCategoryResponse
	result.Transform(category)
	return &result, nil
}

func NewCategoryService(repo *postgresql.Queries) DefaultCategoryService {
	return DefaultCategoryService{repo}
}
