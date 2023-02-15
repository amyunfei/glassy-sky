package service

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/dto"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, data dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, error)
	DeleteCategory(ctx context.Context, data dto.UriIdRequest) error
	ModifyCategory(ctx context.Context, data dto.ModifyCategoryRequest) (*dto.CreateCategoryResponse, error)
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

func (s DefaultCategoryService) DeleteCategory(ctx context.Context, data dto.UriIdRequest) error {
	var err error
	id, err := strconv.ParseInt(data.ID, 10, 64)
	if err != nil {
		return err
	}
	err = s.repo.DeleteCategory(ctx, id)
	return err
}

func (s DefaultCategoryService) ModifyCategory(
	ctx context.Context, data dto.ModifyCategoryRequest,
) (*dto.CreateCategoryResponse, error) {
	id, err := strconv.ParseInt(data.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	name := sql.NullString{}
	if data.Name != "" {
		name.String = data.Name
		name.Valid = true
	}
	parentId := sql.NullInt64{}
	if data.ParentId != "" {
		id, err := strconv.ParseInt(data.ParentId, 10, 64)
		if err != nil {
			return nil, err
		}
		parentId.Int64 = id
		parentId.Valid = true
	}
	color := sql.NullInt32{}
	if data.Color != "" {
		num, err := strconv.ParseInt(data.Color, 16, 32)
		if err != nil {
			return nil, err
		}
		color.Int32 = int32(num)
		color.Valid = true
	}
	arg := postgresql.UpdateCategoryParams{
		ID:       id,
		Name:     name,
		ParentID: parentId,
		Color:    color,
	}
	category, err := s.repo.UpdateCategory(ctx, arg)
	if err != nil {
		return nil, err
	}
	var result dto.CreateCategoryResponse
	result.Transform(category)
	return &result, nil
}

func NewCategoryService(repo *postgresql.Queries) DefaultCategoryService {
	return DefaultCategoryService{repo}
}
