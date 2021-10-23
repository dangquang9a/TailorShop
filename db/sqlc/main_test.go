package db

import (
	"TailorShop/util"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://merlin:@merlin123@10.0.0.223:5432/tailor_shop?sslmode=disable"
// )

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
