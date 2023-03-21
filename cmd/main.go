package main

import "github.com/amyunfei/glassy-sky/internal/admin/app"

// @title Glassy Sky API
// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name authorization
func main() {
	app.Start()
}
