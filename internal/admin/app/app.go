package app

import (
	"database/sql"
	"time"

	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
	"github.com/gin-gonic/gin"
)

func Start() {
	logger.Init()
	defer logger.Sync()

	// db := getDB()
	// queries := postgresql.New(db)
	router := gin.Default()
	router.Run(":9999")
	logger.Info("glassy-sky running on port 9999 ...")
}

func getDB() *sql.DB {
	conn, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/glassy_sky?sslmode=disable")
	if err != nil {
		logger.Panic(err.Error())
	}
	conn.SetConnMaxLifetime(time.Minute * 3) // 设置连接池最大生存时间
	conn.SetMaxOpenConns(10)                 // 设置最大连接池数
	conn.SetMaxIdleConns(10)                 // 设置空闲时的最大连接池数
	return conn
}
