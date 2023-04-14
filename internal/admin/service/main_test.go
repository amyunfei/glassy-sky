package service

import (
	"os"
	"testing"

	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/logger"
)

func TestMain(m *testing.M) {
	logger.Init()
	defer logger.Sync()
	os.Exit(m.Run())
}
