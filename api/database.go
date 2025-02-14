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
	workouts_query := `
		CREATE TABLE IF NOT EXISTS workouts (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	name TEXT NOT NULL,
    	date DATETIME NOT NULL,
    	description TEXT
		);
	`
	db.MustExec(workouts_query)

	activities_query := `
		CREATE TABLE activities (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
  		file_path TEXT
		);
	`
	db.MustExec(activities_query)

	activity_messages_query := `
		CREATE TABLE activity_messages (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	activity_id INTEGER NOT NULL,
  		distance REAL,
  		latitude REAL,
  		longitude REAL,
  		speed REAL,
  		heart_rate INTEGER,
    	FOREIGN KEY (activity_id) REFERENCES Activity(id) ON DELETE CASCADE
		);
	`
	db.MustExec(activity_messages_query)
}
