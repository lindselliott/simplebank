package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testStore *Store
var testQueries *Queries

// var testDB *sql.DB

func TestMain(m *testing.M) {

	conn, err := sql.Open(dbDriver, dbSource)
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
