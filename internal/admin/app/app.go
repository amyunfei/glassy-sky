package app

import (
	_ "github.com/amyunfei/glassy-sky/api"
	"github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql"
	"github.com/amyunfei/glassy-sky/internal/admin/handler"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/database"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
	"github.com/amyunfei/glassy-sky/internal/admin/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start() {
	logger.Init()
	defer logger.Sync()

	db := database.GetDB()
	queries := postgresql.NewStore(db)
	router := gin.Default()

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	categoryHandlers := handler.CategoryHandlers{Service: service.NewCategoryService(queries)}
	router.POST("/category", categoryHandlers.CreateCategory)
	router.DELETE("/category/:id", categoryHandlers.DeleteCategory)
	router.PUT("/category/:id", categoryHandlers.ModifyCategory)
	router.GET("/category", categoryHandlers.ListCategory)

	labelHandlers := handler.LabelHandlers{Service: service.NewLabelService(queries)}
	router.POST("/label", labelHandlers.CreateLabel)
	router.DELETE("/label/:id", labelHandlers.DeleteLabel)
	router.PUT("/label/:id", labelHandlers.ModifyLabel)
	router.GET("/label", labelHandlers.ListLabel)

	router.Run(":9999")
	logger.Info("glassy-sky running on port 9999 ...")
}
