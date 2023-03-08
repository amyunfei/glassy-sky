package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/dto"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/email"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/store"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/utils"
)

type UserService interface {
	SendEmailCode(ctx context.Context, data dto.SendEmailCodeRequest) error
	CheckEmailCode(ctx context.Context, data dto.CheckEmailCodeRequest) (validity bool)
	CreateUser(ctx context.Context, data dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	VerifyEmail(ctx context.Context, data dto.SendEmailCodeRequest) (avaliable bool, err error)
}

type DefaultUserService struct {
	repo postgresql.Repository
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
	arg := postgresql.CreateUserParams{
		Username: data.Username,
		Password: data.Password,
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

func NewUserService(repo postgresql.Repository) UserService {
	return DefaultUserService{repo}
}
