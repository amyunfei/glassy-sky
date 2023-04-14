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

func randomLabel() postgresql.Label {
	return postgresql.Label{
		ID:        utils.RandomInt(0, 99999999),
		Name:      utils.RandomString(6),
		Color:     utils.RandomColorInt(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
	}
}

func TestCreateLabel(t *testing.T) {
	label := randomLabel()
	testCases := []struct {
		name          string
		body          dto.CreateLabelRequest
		buildStubs    func(repo *mockdb.MockRepository)
		checkResponse func(res *dto.CreateLabelResponse, err error)
	}{
		{
			name: "success",
			body: dto.CreateLabelRequest{
				Name:  label.Name,
				Color: strconv.FormatInt(int64(label.Color), 16),
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					CreateLabel(gomock.Any(), gomock.Any()).
					Times(1).
					Return(label, nil)
			},
			checkResponse: func(res *dto.CreateLabelResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, label.Name, res.Name)
				require.Equal(t, strconv.FormatInt(int64(label.Color), 16), res.Color)
			},
		},
		{
			name: "empty name",
			body: dto.CreateLabelRequest{
				Name:  "",
				Color: strconv.FormatInt(int64(label.Color), 16),
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					CreateLabel(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(res *dto.CreateLabelResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
			},
		},
		{
			name: "invalid color",
			body: dto.CreateLabelRequest{
				Name:  label.Name,
				Color: "invalid",
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().
					CreateLabel(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(res *dto.CreateLabelResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockdb.NewMockRepository(ctrl)
			testCase.buildStubs(repo)

			service := NewLabelService(repo)
			testCase.checkResponse(service.CreateLabel(context.Background(), testCase.body))
		})
	}
}

func TestGetLabel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}
