package service

import (
	"os"
	"testing"

	"github.com/amyunfei/glassy-sky/cmd/config"
	"github.com/amyunfei/glassy-sky/internal/admin/app/options"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
)

var (
	testConfig     *config.Config
	testAppOptions *options.AppOptions
)

func TestMain(m *testing.M) {
	logger.Init()
	defer logger.Sync()
	config, err := config.LoadTestConfig("../../../cmd")
	if err != nil {
		logger.Panic(err.Error())
		return
	}
	testConfig = &config
	testAppOptions, err = options.NewAppOptions(&config)
	if err != nil {
		logger.Panic(err.Error())
		return
	}
	os.Exit(m.Run())
}
