package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/util"
)

var testStore Store
var testQueries *Queries

// var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot open db:", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(conn)
	testQueries = New(conn)

	os.Exit(m.Run())
}
