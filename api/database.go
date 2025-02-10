package api

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func SetupDB() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
