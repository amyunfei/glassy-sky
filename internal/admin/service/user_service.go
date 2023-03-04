package service

import (
	"context"

	"github.com/amyunfei/glassy-sky/internal/admin/dto"
)

type UserService interface {
	CreateUser(ctx context.Context, data dto.CreateUserRequest) (*dto.CreateUserResponse, error)
}
