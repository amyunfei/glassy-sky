package service

import (
	"context"
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
	CreateUser(ctx context.Context, data dto.CreateUserRequest) (*dto.CreateUserResponse, error)
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

func (s DefaultUserService) CreateUser(ctx context.Context, data dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	return nil, nil
}

func NewUserService(repo postgresql.Repository) UserService {
	return DefaultUserService{repo}
}
