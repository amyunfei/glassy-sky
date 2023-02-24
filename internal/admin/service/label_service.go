package service

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/dto"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
)

type LabelService interface {
	CreateLabel(ctx context.Context, data dto.CreateLabelRequest) (*dto.CreateLabelResponse, error)
	DeleteLabel(ctx context.Context, data dto.UriIdRequest) error
	ModifyLabel(ctx context.Context, data dto.ModifyLabelRequest) (*dto.CreateLabelResponse, error)
	ListLabel(ctx context.Context, listData dto.ListRequest, filterData dto.FilterLabelRequest) (*dto.ListResponse[dto.CreateLabelResponse], error)
}

type DefaultLabelService struct {
	repo *postgresql.Queries
}

func (s DefaultLabelService) CreateLabel(
	ctx context.Context, data dto.CreateLabelRequest,
) (*dto.CreateLabelResponse, error) {
	color, err := strconv.ParseInt(data.Color, 16, 32)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	arg := postgresql.CreateLabelParams{
		Name:  data.Name,
		Color: int32(color),
	}
	label, err := s.repo.CreateLabel(ctx, arg)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	var result dto.CreateLabelResponse
	result.Transform(label)
	return &result, nil
}

func (s DefaultLabelService) DeleteLabel(ctx context.Context, data dto.UriIdRequest) error {
	var err error
	id, err := strconv.ParseInt(data.ID, 10, 64)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = s.repo.DeleteLabel(ctx, id)
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

func (s DefaultLabelService) ModifyLabel(
	ctx context.Context, data dto.ModifyLabelRequest,
) (*dto.CreateLabelResponse, error) {
	id, err := strconv.ParseInt(data.ID, 10, 64)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	name := sql.NullString{}
	if data.Name != "" {
		name.String = data.Name
		name.Valid = true
	}
	color := sql.NullInt32{}
	if data.Color != "" {
		num, err := strconv.ParseInt(data.Color, 16, 32)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		color.Int32 = int32(num)
		color.Valid = true
	}
	arg := postgresql.UpdateLabelParams{
		ID:    id,
		Name:  name,
		Color: color,
	}
	label, err := s.repo.UpdateLabel(ctx, arg)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	var result dto.CreateLabelResponse
	result.Transform(label)
	return &result, nil
}

func (s DefaultLabelService) ListLabel(
	ctx context.Context, listData dto.ListRequest, filterData dto.FilterLabelRequest,
) (*dto.ListResponse[dto.CreateLabelResponse], error) {
	arg := postgresql.ListLabelParams{
		Limit:  listData.Size,
		Offset: (listData.Current - 1) * listData.Size,
		Name:   filterData.Name,
	}
	count, err := s.repo.CountLabel(ctx, filterData.Name)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	labels, err := s.repo.ListLabel(ctx, arg)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	list := make([]dto.CreateLabelResponse, 0)
	for _, label := range labels {
		var item dto.CreateLabelResponse
		item.Transform(label)
		list = append(list, item)
	}
	return &dto.ListResponse[dto.CreateLabelResponse]{
		List:  list,
		Count: count,
	}, nil
}

func NewLabelService(repo *postgresql.Queries) LabelService {
	return DefaultLabelService{repo}
}