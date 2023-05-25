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

func randomUser() postgresql.User {
	return postgresql.User{
		ID:       utils.RandomInt(0, 99999999),
		Username: utils.RandomString(6),
		Password: utils.RandomString(16),
		Email:    utils.RandomEmail(),
		Nickname: utils.RandomString(6),
		Avatar: sql.NullString{
			String: "",
			Valid:  false,
		},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
	}
}

func TestCreateUser(t *testing.T) {
	user := randomUser()
	testCases := []struct {
		name          string
		body          dto.CreateUserRequest
		buildStubs    func(repo *mockdb.MockRepository)
		checkResponse func(res *dto.CreateUserResponse, err error)
	}{
		{
			name: "success",
			body: dto.CreateUserRequest{
				Username: user.Username,
				Password: user.Password,
				Email:    user.Email,
				Nickname: user.Nickname,
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(res *dto.CreateUserResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, user.Username, res.Username)
				require.Equal(t, user.Email, res.Email)
				require.Equal(t, user.Nickname, res.Nickname)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockdb.NewMockRepository(ctrl)
			testCase.buildStubs(repo)

			service := NewUserService(repo, nil)
			testCase.checkResponse(service.CreateUser(context.Background(), testCase.body))
		})
	}
}

func TestCreateSuperAdmin(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(repo *mockdb.MockRepository)
		checkResponse func(err error)
	}{
		{
			name: "super_admin_exists",
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					CountUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(int64(1), nil)
				repo.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "create_super_admin_when_no__exists",
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					CountUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(int64(0), nil)
				repo.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1)
			},
			checkResponse: func(err error) {
				require.NoError(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockdb.NewMockRepository(ctrl)
			testCase.buildStubs(repo)

			service := NewUserService(repo, nil)
			testCase.checkResponse(service.CreateSuperAdmin(context.Background()))
		})
	}
}

func TestVerifyEmail(t *testing.T) {
	user := randomUser()
	testCases := []struct {
		name          string
		body          dto.SendEmailCodeRequest
		buildStubs    func(repo *mockdb.MockRepository)
		checkResponse func(avaliable bool, err error)
	}{
		{
			name: "email_exists",
			body: dto.SendEmailCodeRequest{
				Email: user.Email,
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(avaliable bool, err error) {
				require.NoError(t, err)
				require.Equal(t, false, avaliable)
			},
		},
		{
			name: "email_not_exists",
			body: dto.SendEmailCodeRequest{
				Email: "not_exists_email",
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(postgresql.User{}, sql.ErrNoRows)
			},
			checkResponse: func(avaliable bool, err error) {
				require.NoError(t, err)
				require.Equal(t, true, avaliable)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockdb.NewMockRepository(ctrl)
			testCase.buildStubs(repo)

			service := NewUserService(repo, nil)
			testCase.checkResponse(service.VerifyEmail(context.Background(), testCase.body))
		})
	}
}

func TestModifyUser(t *testing.T) {
	user := randomUser()
	testCases := []struct {
		name          string
		body          dto.ModifyUserRequest
		buildStubs    func(repo *mockdb.MockRepository)
		checkResponse func(res *dto.CreateUserResponse, err error)
	}{
		{
			name: "success_modify_user",
			body: dto.ModifyUserRequest{
				ID:       strconv.FormatInt(user.ID, 10),
				Password: user.Password,
				Nickname: user.Nickname,
				Avatar:   utils.RandomString(10),
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(res *dto.CreateUserResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, user.Username, res.Username)
				require.Equal(t, user.Email, res.Email)
				require.Equal(t, user.Nickname, res.Nickname)
			},
		},
		{
			name: "fail_modify_user_with_invalid_id",
			body: dto.ModifyUserRequest{
				ID:       "invalid",
				Password: user.Password,
				Nickname: user.Nickname,
				Avatar:   "",
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(res *dto.CreateUserResponse, err error) {
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

			service := NewUserService(repo, nil)
			testCase.checkResponse(service.ModifyUser(context.Background(), testCase.body))
		})
	}
}

func TestListUser(t *testing.T) {
	users := make([]postgresql.User, 0)
	for i := 0; i < 10; i++ {
		users = append(users, randomUser())
	}

	testCases := []struct {
		name          string
		body          dto.ListRequest
		filterData    dto.FilterUserRequest
		buildStubs    func(repo *mockdb.MockRepository)
		checkResponse func(res *dto.ListResponse[dto.CreateUserResponse], err error)
	}{
		{
			name: "success_list_category",
			body: dto.ListRequest{
				Size:    10,
				Current: 1,
			},
			filterData: dto.FilterUserRequest{
				Username: "",
				Nickname: "",
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					ExecTx(gomock.Any(), gomock.Any()).
					Do(func(_ context.Context, fn func(postgresql.Querier) error) {
						repo.EXPECT().
							CountUser(gomock.Any(), gomock.Any()).
							Times(1).
							Return(int64(len(users)), nil)

						repo.EXPECT().
							ListUser(gomock.Any(), gomock.Any()).
							Times(1).
							Return(users, nil)
						fn(repo)
					}).Return(nil)
			},
			checkResponse: func(res *dto.ListResponse[dto.CreateUserResponse], err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, int64(len(users)), res.Count)
				require.Len(t, res.List, len(users))
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockdb.NewMockRepository(ctrl)
			testCase.buildStubs(repo)

			service := NewUserService(repo, nil)
			testCase.checkResponse(service.ListUser(context.Background(), testCase.body, testCase.filterData))
		})
	}
}
