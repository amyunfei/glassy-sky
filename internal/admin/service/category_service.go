package service

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/dto"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/utils"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, data dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, error)
	DeleteCategory(ctx context.Context, data dto.UriIdRequest) error
	ModifyCategory(ctx context.Context, data dto.ModifyCategoryRequest) (*dto.CreateCategoryResponse, error)
	ListCategory(ctx context.Context, listData dto.ListRequest, filterData dto.FilterCategoryRequest) (*dto.ListResponse[dto.CreateCategoryResponse], error)
	GetCategory(ctx context.Context, data dto.UriIdRequest) (*dto.CreateCategoryResponse, error)
}

type DefaultCategoryService struct {
	repo postgresql.Repository
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
	color, err := utils.HexColorToInt[int32](data.Color)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	arg := postgresql.CreateCategoryParams{
		Name:     data.Name,
		ParentID: parentId,
		Color:    color,
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
		num, err := utils.HexColorToInt[int32](data.Color)
		if err != nil {
			return nil, err
		}
		color.Int32 = num
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

func (s DefaultCategoryService) ListCategory(
	ctx context.Context, listData dto.ListRequest, filterData dto.FilterCategoryRequest,
) (*dto.ListResponse[dto.CreateCategoryResponse], error) {
	var result dto.ListResponse[dto.CreateCategoryResponse]
	err := s.repo.ExecTx(ctx, func(q postgresql.Querier) error {
		var err error
		arg := postgresql.ListCategoryParams{
			Limit:  listData.Size,
			Offset: (listData.Current - 1) * listData.Size,
			Name:   filterData.Name,
		}
		count, err := q.CountCategory(ctx, filterData.Name)
		if err != nil {
			return err
		}
		categories, err := q.ListCategory(ctx, arg)
		if err != nil {
			return err
		}
		list := make([]dto.CreateCategoryResponse, 0)
		for _, category := range categories {
			var item dto.CreateCategoryResponse
			item.Transform(category)
			list = append(list, item)
		}
		result.Count = count
		result.List = list
		return nil
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return &result, nil
}

func (s DefaultCategoryService) GetCategory(
	ctx context.Context, data dto.UriIdRequest,
) (*dto.CreateCategoryResponse, error) {
	id, err := strconv.ParseInt(data.ID, 10, 64)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	category, err := s.repo.GetCategory(ctx, id)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	var result dto.CreateCategoryResponse
	result.Transform(category)
	return &result, nil
}

func NewCategoryService(repo postgresql.Repository) CategoryService {
	return DefaultCategoryService{repo}
}
