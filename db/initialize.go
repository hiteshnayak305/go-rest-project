package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var SqlConnection *sql.DB

func Initialize() {
	var err error
	SqlConnection, err = sql.Open("sqlite3", "events.db")

	if err != nil {
		panic(err)
	}

	err = SqlConnection.Ping()
	if err != nil {
		panic(err)
	}

	SqlConnection.SetMaxOpenConns(10)
	SqlConnection.SetMaxIdleConns(5)

	inititlizeTables()
}

func inititlizeTables() {

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER
	)`

	_, err := SqlConnection.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}
}
