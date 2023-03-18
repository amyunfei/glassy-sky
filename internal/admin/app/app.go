package app

import (
	_ "github.com/amyunfei/glassy-sky/api"
	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/database"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/token"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start() {
	logger.Init()
	defer logger.Sync()

	db := database.GetDB()
	queries := postgresql.NewStore(db)
	tokenMaker, err := token.NewJWTMaker("secret")
	if err != nil {
		logger.Panic(err.Error())
		return
	}
	router := gin.Default()

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userHandlers := InitializeUserHandlers(queries, tokenMaker)
	router.GET("/user/email-verify/:email", userHandlers.VerifyEmail)
	router.GET("/user/email-code/:email", userHandlers.SendEmailCode)
	router.POST("/user/register", userHandlers.RegisterUser)
	router.POST("/user/login", userHandlers.Login)

	router.GET("/user", userHandlers.ListUser)
	router.PUT("/user/:id", userHandlers.ModifyUser)

	categoryHandlers := InitializeCategoryHandlers(queries)
	router.POST("/category", categoryHandlers.CreateCategory)
	router.DELETE("/category/:id", categoryHandlers.DeleteCategory)
	router.PUT("/category/:id", categoryHandlers.ModifyCategory)
	router.GET("/category", categoryHandlers.ListCategory)

	labelHandlers := InitializeLabelHandlers(queries)
	router.POST("/label", labelHandlers.CreateLabel)
	router.DELETE("/label/:id", labelHandlers.DeleteLabel)
	router.PUT("/label/:id", labelHandlers.ModifyLabel)
	router.GET("/label", labelHandlers.ListLabel)

	router.Run(":9999")
	logger.Info("glassy-sky running on port 9999 ...")
}
