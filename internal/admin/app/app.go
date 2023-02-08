package app

import (
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
	"github.com/gin-gonic/gin"
)

func Start() {
	logger.Init()
	defer logger.Sync()

	router := gin.Default()
	router.Run(":9999")
	logger.Info("glassy-sky running on port 9999 ...")
}
