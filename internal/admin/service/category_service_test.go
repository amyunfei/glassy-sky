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

func randomCategory() postgresql.Category {
	return postgresql.Category{
		ID:        utils.RandomInt(0, 99999999),
		Name:      utils.RandomString(6),
		ParentID:  0,
		Color:     utils.RandomColorInt(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
	}
}

func TestCreateCategory(t *testing.T) {
	parentCategory := randomCategory()
	category := randomCategory()
	category.ParentID = parentCategory.ID

	testCases := []struct {
		name          string
		body          dto.CreateCategoryRequest
		buildStubs    func(*mockdb.MockRepository)
		checkResponse func(*dto.CreateCategoryResponse, error)
	}{
		{
			name: "success_create_root_category",
			body: dto.CreateCategoryRequest{
				Name:     category.Name,
				Color:    utils.IntToHexColor(category.Color),
				ParentId: "",
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().GetCategory(gomock.Any(), gomock.Any()).Times(0)
				repo.EXPECT().
					CreateCategory(gomock.Any(), gomock.Any()).
					Times(1).
					Return(category, nil)
			},
			checkResponse: func(res *dto.CreateCategoryResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, category.Name, res.Name)
				require.Equal(t, utils.IntToHexColor(category.Color), res.Color)
			},
		},
		{
			name: "success_create_sub_category",
			body: dto.CreateCategoryRequest{
				Name:     category.Name,
				Color:    utils.IntToHexColor(category.Color),
				ParentId: strconv.FormatInt(category.ParentID, 10),
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					GetCategory(gomock.Any(), gomock.Eq(category.ParentID)).
					Times(1).
					Return(parentCategory, nil)
				repo.EXPECT().
					CreateCategory(gomock.Any(), gomock.Any()).
					Times(1).
					Return(category, nil)
			},
			checkResponse: func(res *dto.CreateCategoryResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, category.Name, res.Name)
				require.Equal(t, utils.IntToHexColor(category.Color), res.Color)
			},
		},
		{
			name: "fail_create_sub_category_with_invalid_parent_id",
			body: dto.CreateCategoryRequest{
				Name:     category.Name,
				Color:    utils.IntToHexColor(category.Color),
				ParentId: "invalid",
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					GetCategory(gomock.Any(), gomock.Eq(category.ParentID)).
					Times(0)
			},
			checkResponse: func(res *dto.CreateCategoryResponse, err error) {
				_, err1 := strconv.ParseInt("invalid", 10, 64)
				require.Error(t, err)
				require.EqualError(t, err, err1.Error())
			},
		},
		{
			name: "fail_create_sub_category_with_no-existed_parent_id",
			body: dto.CreateCategoryRequest{
				Name:     category.Name,
				Color:    utils.IntToHexColor(category.Color),
				ParentId: strconv.FormatInt(category.ParentID, 10),
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					GetCategory(gomock.Any(), gomock.Eq(category.ParentID)).
					Times(1).
					Return(postgresql.Category{}, sql.ErrNoRows)
			},
			checkResponse: func(res *dto.CreateCategoryResponse, err error) {
				require.Error(t, err)
				require.Equal(t, err, sql.ErrNoRows)
			},
		},
		{
			name: "fail_create_sub_category_with_invalid_color",
			body: dto.CreateCategoryRequest{
				Name:     category.Name,
				Color:    "invalid",
				ParentId: "",
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().GetCategory(gomock.Any(), gomock.Any()).Times(0)
				repo.EXPECT().CreateCategory(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(res *dto.CreateCategoryResponse, err error) {
				_, err1 := utils.HexColorToInt[int32]("invalid")
				require.Error(t, err)
				require.EqualError(t, err, err1.Error())
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockdb.NewMockRepository(ctrl)
			testCase.buildStubs(repo)

			service := NewCategoryService(repo)
			testCase.checkResponse(service.CreateCategory(context.Background(), testCase.body))
		})
	}
}

func TestGetCategory(t *testing.T) {
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

func TestListCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}
