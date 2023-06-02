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

			service := NewCategoryService(repo, testAppOptions)
			testCase.checkResponse(service.CreateCategory(context.Background(), testCase.body))
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	category := randomCategory()
	testCases := []struct {
		name          string
		body          dto.UriIdRequest
		buidStubs     func(*mockdb.MockRepository)
		checkResponse func(error)
	}{
		{
			name: "success_delete_category",
			body: dto.UriIdRequest{
				ID: strconv.FormatInt(category.ID, 10),
			},
			buidStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					DeleteCategory(gomock.Any(), gomock.Eq(category.ID)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "fail_delete_category_with_invalid_id",
			body: dto.UriIdRequest{
				ID: "invalid",
			},
			buidStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().DeleteCategory(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(err error) {
				_, err1 := strconv.ParseInt("invalid", 10, 64)
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
			testCase.buidStubs(repo)

			service := NewCategoryService(repo, testAppOptions)
			testCase.checkResponse(service.DeleteCategory(context.Background(), testCase.body))
		})
	}
}

func TestModifyCategory(t *testing.T) {
	category := randomCategory()
	testCases := []struct {
		name          string
		body          dto.ModifyCategoryRequest
		buildStubs    func(repo *mockdb.MockRepository)
		checkResponse func(res *dto.CreateCategoryResponse, err error)
	}{
		{
			name: "success_modify_category_with_empty_parent_id",
			body: dto.ModifyCategoryRequest{
				ID:       strconv.FormatInt(category.ID, 10),
				Name:     category.Name,
				Color:    utils.IntToHexColor(category.Color),
				ParentId: "",
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					UpdateCategory(gomock.Any(), gomock.Any()).
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
			name: "success_modify_category_with_parent_id",
			body: dto.ModifyCategoryRequest{
				ID:       strconv.FormatInt(category.ID, 10),
				Name:     category.Name,
				Color:    utils.IntToHexColor(category.Color),
				ParentId: strconv.FormatInt(1, 10),
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					UpdateCategory(gomock.Any(), gomock.Any()).
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
			name: "fail_modify_category_with_invalid_parent_id",
			body: dto.ModifyCategoryRequest{
				ID:       strconv.FormatInt(category.ID, 10),
				Name:     category.Name,
				Color:    utils.IntToHexColor(category.Color),
				ParentId: "invalid",
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().UpdateCategory(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(res *dto.CreateCategoryResponse, err error) {
				_, err1 := strconv.ParseInt("invalid", 10, 64)
				require.Error(t, err)
				require.EqualError(t, err, err1.Error())
			},
		},
		{
			name: "fail_modify_category_with_invalid_id",
			body: dto.ModifyCategoryRequest{
				ID:       "invalid",
				Name:     category.Name,
				Color:    utils.IntToHexColor(category.Color),
				ParentId: "",
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().UpdateCategory(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(res *dto.CreateCategoryResponse, err error) {
				_, err1 := strconv.ParseInt("invalid", 10, 64)
				require.Error(t, err)
				require.EqualError(t, err, err1.Error())
			},
		},
		{
			name: "fail_modify_category_with_invalid_color",
			body: dto.ModifyCategoryRequest{
				ID:       strconv.FormatInt(category.ID, 10),
				Name:     category.Name,
				Color:    "invalid",
				ParentId: "",
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().UpdateCategory(gomock.Any(), gomock.Any()).Times(0)
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

			service := NewCategoryService(repo, testAppOptions)
			testCase.checkResponse(service.ModifyCategory(context.Background(), testCase.body))
		})
	}
}

func TestGetCategory(t *testing.T) {
	category := randomCategory()
	testCases := []struct {
		name          string
		body          dto.UriIdRequest
		buildStubs    func(repo *mockdb.MockRepository)
		checkResponse func(res *dto.CreateCategoryResponse, err error)
	}{
		{
			name: "success_get_category",
			body: dto.UriIdRequest{
				ID: strconv.FormatInt(category.ID, 10),
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					GetCategory(gomock.Any(), gomock.Eq(category.ID)).
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
			name: "fail_get_category_with_invalid_id",
			body: dto.UriIdRequest{
				ID: "invalid",
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().GetCategory(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(res *dto.CreateCategoryResponse, err error) {
				_, err1 := strconv.ParseInt("invalid", 10, 64)
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

			service := NewCategoryService(repo, testAppOptions)
			testCase.checkResponse(service.GetCategory(context.Background(), testCase.body))
		})
	}
}

func TestListCategory(t *testing.T) {
	categories := make([]postgresql.Category, 0)
	for i := 0; i < 10; i++ {
		categories = append(categories, randomCategory())
	}

	testCases := []struct {
		name          string
		body          dto.ListRequest
		filterData    dto.FilterCategoryRequest
		buildStubs    func(repo *mockdb.MockRepository)
		checkResponse func(res *dto.ListResponse[dto.CreateCategoryResponse], err error)
	}{
		{
			name: "success_list_category",
			body: dto.ListRequest{
				Size:    10,
				Current: 1,
			},
			filterData: dto.FilterCategoryRequest{
				Name: "",
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					ExecTx(gomock.Any(), gomock.Any()).
					Do(func(_ context.Context, fn func(postgresql.Querier) error) {
						repo.EXPECT().
							CountCategory(gomock.Any(), gomock.Any()).
							Times(1).
							Return(int64(len(categories)), nil)

						repo.EXPECT().
							ListCategory(gomock.Any(), gomock.Any()).
							Times(1).
							Return(categories, nil)
						fn(repo)
					}).Return(nil)
			},
			checkResponse: func(res *dto.ListResponse[dto.CreateCategoryResponse], err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, int64(len(categories)), res.Count)
				require.Len(t, res.List, len(categories))
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockdb.NewMockRepository(ctrl)
			testCase.buildStubs(repo)

			service := NewCategoryService(repo, testAppOptions)
			testCase.checkResponse(service.ListCategory(context.Background(), testCase.body, testCase.filterData))
		})
	}
}
