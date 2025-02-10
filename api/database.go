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

	createTables(db)
	return db
}

func createTables(db *sqlx.DB) {
	query := `
		CREATE TABLE my_table (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	name TEXT NOT NULL,
    	date DATETIME NOT NULL,
    	description TEXT
		);
	`
	db.MustExec(query)
}
