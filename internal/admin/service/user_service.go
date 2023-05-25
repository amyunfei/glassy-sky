package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/dto"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/email"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/store"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/token"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/utils"
)

type UserService interface {
	SendEmailCode(ctx context.Context, data dto.SendEmailCodeRequest) error
	CheckEmailCode(ctx context.Context, data dto.CheckEmailCodeRequest) (validity bool)
	CreateUser(ctx context.Context, data dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	CreateSuperAdmin(ctx context.Context) error
	VerifyEmail(ctx context.Context, data dto.SendEmailCodeRequest) (avaliable bool, err error)
	Login(ctx context.Context, data dto.LoginRequest) (token string, err error)
	ListUser(
		ctx context.Context, listData dto.ListRequest, filterData dto.FilterUserRequest,
	) (*dto.ListResponse[dto.CreateUserResponse], error)
	ModifyUser(
		ctx context.Context, data dto.ModifyUserRequest,
	) (*dto.CreateUserResponse, error)
}

type DefaultUserService struct {
	repo       postgresql.Repository
	tokenMaker token.Maker
}

func (s DefaultUserService) SendEmailCode(ctx context.Context, data dto.SendEmailCodeRequest) error {
	code := utils.RandomCode(6)
	err := email.SendEmail(data.Email, "验证码发送", "您的验证码: <b>"+code+"</b>")
	if err != nil {
		logger.Error(err.Error())
	} else {
		store.Set(data.Email, code, time.Minute)
	}
	return err
}

func (s DefaultUserService) CheckEmailCode(ctx context.Context, data dto.CheckEmailCodeRequest) (validity bool) {
	code, ok := store.KV.Load(data.Email)
	if ok && code == data.Code {
		return true
	} else {
		return false
	}
}

func (s DefaultUserService) CreateUser(ctx context.Context, data dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	password, err := utils.HashPassword(data.Password)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	arg := postgresql.CreateUserParams{
		Username: data.Username,
		Password: password,
		Email:    data.Email,
		Nickname: data.Nickname,
	}
	user, err := s.repo.CreateUser(ctx, arg)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	var result dto.CreateUserResponse
	result.Transform(user)
	return &result, nil
}

func (s DefaultUserService) CreateSuperAdmin(ctx context.Context) error {
	count, err := s.repo.CountUser(ctx, postgresql.CountUserParams{
		Username: "admin",
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if count != 0 {
		return nil
	}
	password, err := utils.HashPassword("a123456")
	fmt.Println(password, err)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	arg := postgresql.CreateUserParams{
		Username: "admin",
		Password: password,
		Email:    "amyunfei@163.com",
		Nickname: "超级管理员",
	}
	user, err := s.repo.CreateUser(ctx, arg)
	fmt.Println(user, err)
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

func (s DefaultUserService) VerifyEmail(ctx context.Context, data dto.SendEmailCodeRequest) (avaliable bool, err error) {
	_, err = s.repo.GetUserByEmail(ctx, data.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			avaliable = true
			err = nil
		} else {
			logger.Error(err.Error())
			avaliable = false
		}
	} else {
		avaliable = false
	}
	return
}

func (s DefaultUserService) Login(ctx context.Context, data dto.LoginRequest) (token string, err error) {
	user, err := s.repo.GetUserByUsername(ctx, data.Username)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("username or password incorrect")
	}
	err = utils.CheckPassword(data.Password, user.Password)
	if err != nil {
		return "", errors.New("username or password incorrect")
	}
	token, err = s.tokenMaker.CreateToken(user.Username, time.Minute)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return token, nil
}

func (s DefaultUserService) ModifyUser(
	ctx context.Context, data dto.ModifyUserRequest,
) (*dto.CreateUserResponse, error) {
	id, err := strconv.ParseInt(data.ID, 10, 64)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	nickname := sql.NullString{}
	if data.Nickname != "" {
		nickname.String = data.Nickname
		nickname.Valid = true
	}
	avatar := sql.NullString{}
	if data.Avatar != "" {
		avatar.String = data.Avatar
		avatar.Valid = true
	}
	arg := postgresql.UpdateUserParams{
		ID:       id,
		Nickname: nickname,
		Avatar:   avatar,
		Password: sql.NullString{
			String: "",
			Valid:  false,
		},
	}
	user, err := s.repo.UpdateUser(ctx, arg)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	var result dto.CreateUserResponse
	result.Transform(user)
	return &result, nil
}

func (s DefaultUserService) ListUser(
	ctx context.Context,
	listData dto.ListRequest,
	filterData dto.FilterUserRequest,
) (*dto.ListResponse[dto.CreateUserResponse], error) {
	var result dto.ListResponse[dto.CreateUserResponse]
	err := s.repo.ExecTx(ctx, func(q postgresql.Querier) error {
		arg := postgresql.ListUserParams{
			Limit:    listData.Size,
			Offset:   (listData.Current - 1) * listData.Size,
			Username: filterData.Username,
			Nickname: filterData.Nickname,
		}
		count, err := s.repo.CountUser(ctx, postgresql.CountUserParams{
			Username: filterData.Username,
			Nickname: filterData.Nickname,
		})
		if err != nil {
			return err
		}
		users, err := s.repo.ListUser(ctx, arg)
		if err != nil {
			return err
		}
		list := make([]dto.CreateUserResponse, 0)
		for _, user := range users {
			var item dto.CreateUserResponse
			item.Transform(user)
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

func NewUserService(repo postgresql.Repository, tokenMaker token.Maker) UserService {
	return DefaultUserService{repo, tokenMaker}
}
