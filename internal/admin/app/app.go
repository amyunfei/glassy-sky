package app

import (
	"context"
	"fmt"

	_ "github.com/amyunfei/glassy-sky/api"
	"github.com/amyunfei/glassy-sky/cmd/config"
	"github.com/amyunfei/glassy-sky/internal/admin/app/options"
	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/database"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
	"github.com/amyunfei/glassy-sky/internal/admin/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start(config *config.Config) {
	logger.Init()
	defer logger.Sync()

	db := database.GetDB(config.DBDriver, config.DBSource)
	queries := postgresql.NewStore(db)

	appOptions, err := options.NewAppOptions(config)
	if err != nil {
		logger.Panic(err.Error())
		return
	}
	router := gin.Default()

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userHandlers := InitializeUserHandlers(queries, appOptions.GetTokenMaker(), config)
	userHandlers.Service.CreateSuperAdmin(context.Background())
	router.GET("/user/email-verify/:email", userHandlers.VerifyEmail)
	router.GET("/user/email-code/:email", userHandlers.SendEmailCode)
	router.POST("/user/register", userHandlers.RegisterUser)
	router.POST("/user/login", userHandlers.Login)

	authRouter := router.Group("/")
	authRouter.Use(middleware.AuthMiddleware(appOptions.GetTokenMaker()))
	authRouter.POST("/user", userHandlers.CreateUser)
	authRouter.GET("/user", userHandlers.ListUser)
	authRouter.GET("/user/:id", userHandlers.GetUser)
	authRouter.PUT("/user/:id", userHandlers.ModifyUser)
	authRouter.DELETE("/user/:id", userHandlers.DeleteUser)

	categoryHandlers := InitializeCategoryHandlers(queries, appOptions)
	authRouter.POST("/category", categoryHandlers.CreateCategory)
	authRouter.DELETE("/category/:id", categoryHandlers.DeleteCategory)
	authRouter.PUT("/category/:id", categoryHandlers.ModifyCategory)
	authRouter.GET("/category", categoryHandlers.ListCategory)

	labelHandlers := InitializeLabelHandlers(queries, appOptions)
	authRouter.POST("/label", labelHandlers.CreateLabel)
	authRouter.DELETE("/label/:id", labelHandlers.DeleteLabel)
	authRouter.PUT("/label/:id", labelHandlers.ModifyLabel)
	authRouter.GET("/label", labelHandlers.ListLabel)
	authRouter.GET("/label/:id", labelHandlers.GetLabel)

	router.Run(config.ServerAddress)
	logger.Info(fmt.Sprintf("glassy-sky running on %s ...", config.ServerAddress))
}
