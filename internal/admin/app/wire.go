//go:build wireinject
// +build wireinject

package app

import (
	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/handler"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/token"
	"github.com/amyunfei/glassy-sky/internal/admin/service"
	"github.com/google/wire"
)

// user handlers
func InitializeUserHandlers(repo postgresql.Repository, tokenMaker token.Maker) handler.UserHandlers {
	wire.Build(
		handler.NewUserHandlers,
		service.NewUserService,
	)
	return handler.UserHandlers{}
}

// category handlers
func InitializeCategoryHandlers(repo postgresql.Repository) handler.CategoryHandlers {
	wire.Build(
		handler.NewCategoryHandlers,
		service.NewCategoryService,
	)
	return handler.CategoryHandlers{}
}

// label handlers
func InitializeLabelHandlers(repo postgresql.Repository) handler.LabelHandlers {
	wire.Build(
		handler.NewLabelHandlers,
		service.NewLabelService,
	)
	return handler.LabelHandlers{}
}
