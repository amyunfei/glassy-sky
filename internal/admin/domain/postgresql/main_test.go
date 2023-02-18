package postgresql

import (
	"database/sql"
	"os"
	"testing"

	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/database"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	testDB = database.GetDB()
	testQueries = New(testDB)
	os.Exit(m.Run())
}
