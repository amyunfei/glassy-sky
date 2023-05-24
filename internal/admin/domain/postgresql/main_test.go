package postgresql

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/amyunfei/glassy-sky/cmd/config"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/database"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	testConfig, err := config.LoadTestConfig("../../../../cmd")
	if err != nil {
		log.Fatal("cannot load test config:", err)
	}
	testDB = database.GetDB(testConfig.DBDriver, testConfig.DBSource)
	testQueries = New(testDB)
	os.Exit(m.Run())
}
