package service

import (
	"os"
	"testing"

	"github.com/amyunfei/glassy-sky/cmd/config"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
)

var testConfig *config.Config

func TestMain(m *testing.M) {
	logger.Init()
	defer logger.Sync()
	config, err := config.LoadTestConfig("../../../cmd")
	if err != nil {
		logger.Panic(err.Error())
	}
	testConfig = &config
	os.Exit(m.Run())
}
