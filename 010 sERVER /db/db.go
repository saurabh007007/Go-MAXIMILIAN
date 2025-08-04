package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqllite3", "api.db")

	if err != nil {
		panic("Could not connect DB")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	CreateTables()

}
func CreateTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	datetime DATETIME NOT NULL,
	user_id INTEGER
	)	
	`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Create not tabele")
	}
}
