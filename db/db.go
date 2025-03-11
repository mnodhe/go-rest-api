package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")
	if err != nil {
		panic("Could not connect to database.")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	CreateTables()
}
func CreateTables() {
	createUsersTable := `
	create table if not exists users (
		id integer primary key autoincrement,
		email text not null unique,
		password text not null
)`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("could not create users table")
	}
	createEventTable := `
	CREATE TABLE IF NOT EXISTS event (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		foreign key (user_id) REFERENCES users(id)
	)
`
	_, err = DB.Exec(createEventTable)
	if err != nil {
		panic("Could not create event table.")
	}
}
