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

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := SqlConnection.Exec(createUsersTable)
	if err != nil {
		panic(err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = SqlConnection.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}
}
