package tests

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/aliforever/go-cad"
	_ "github.com/lib/pq"
)

func TestValuerScanner(t *testing.T) {
	const (
		dbHost     = "localhost"
		dbPort     = 5432
		dbName     = "gocad_test"
		dbUser     = "postgres"
		dbPassword = "159852"
	)

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName))

	if err != nil {
		t.Fatal("Test Failed.", err)
		return
	}

	defer db.Close()

	c := cad.Cents(452)
	_, err = db.Exec(
		"insert into users (name, balance) values ($1, $2)",
		"Ali", c,
	)

	if err != nil {
		t.Fatal("Test Failed.", err)
		return
	}

	var (
		name    string
		balance cad.CAD
	)

	row := db.QueryRow("select name, balance from users where name = $1", "Ali")
	err = row.Scan(&name, &balance)
	if err != nil {
		t.Fatal("Test Failed.", err)
		return
	}

	if balance.AsCents() != 452 {
		t.Fatal(fmt.Scan("Test Failed\nExpected: %d\nResult: %d", 452, balance.AsCents()))
		return
	}
}
