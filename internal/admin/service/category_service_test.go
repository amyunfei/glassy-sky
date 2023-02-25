package service

import (
	"context"
	"database/sql"
	"strconv"
	"testing"
	"time"

	"github.com/amyunfei/glassy-sky/internal/admin/domain/mockdb"
	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/dto"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateCategory(t *testing.T) {
	category := randomCategory()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// build stubs
	repo := mockdb.NewMockRepository(ctrl)
	repo.EXPECT().
		GetCategory(gomock.Any(), gomock.Eq(category.ID)).
		Times(1).
		Return(category, nil)

	service := NewCategoryService(repo)
	arg := dto.UriIdRequest{
		ID: strconv.FormatInt(category.ID, 10),
	}
	var res1 dto.CreateCategoryResponse
	res1.Transform(category)
	res2, err := service.GetCategory(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, res2)
	require.Equal(t, res1, *res2)
}

func randomCategory() postgresql.Category {
	return postgresql.Category{
		ID:        utils.RandomInt(0, 99999999),
		Name:      utils.RandomString(6),
		ParentID:  0,
		Color:     0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
	}
}
