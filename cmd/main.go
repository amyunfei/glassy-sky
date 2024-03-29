package main

import (
	"log"

	"github.com/amyunfei/glassy-sky/cmd/config"
	"github.com/amyunfei/glassy-sky/internal/admin/app"
)

// @title Glassy Sky API
// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name authorization
func main() {
	config, err := config.LoadConfig("./cmd")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	app.Start(&config)
}
