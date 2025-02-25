package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var db *sql.DB

func SetDb() error {
	var err error
	db, err = sql.Open("libsql", fmt.Sprintf("%s?authToken=%s", os.Getenv("DB_URL"), os.Getenv("DB_TOKEN")))

	if err != nil {
		return err
	}

	return db.Ping()
}
