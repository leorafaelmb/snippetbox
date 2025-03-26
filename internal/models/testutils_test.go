package models

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func newTestDB(t *testing.T) *sql.DB {
	dbCreds := os.Getenv("SB_TDB_CREDS")
	db, err := sql.Open("mysql", fmt.Sprintf("%s@/test_snippetbox?parseTime=true&multiStatements=true", dbCreds))
	if err != nil {
		t.Fatal(err)
	}

	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		db.Close()
		t.Fatal(err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		db.Close()
		t.Fatal(err)
	}

	t.Cleanup(func() {
		defer db.Close()

		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}

		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}
	})

	return db
}
