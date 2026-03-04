package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("could not connect to database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	CreateTables()
}
func CreateTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id integer primary key autoincrement,
		name text not null,
		description text not null,
		location text not null,
		dateTime datetime not null,
		user_id integer
	)
	`
	//` ` helps us to write string in multiple line
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id integer primary key autoincrement,
		email text not null unique,
		password text not null
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("could not create users table")
	}
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("could not create events table")
	}

	createRegistrationsTable := `
	create table if not exists registrations (
	id integer primary key autoincrement,
	event_id integer,
	user_id integer,
	foreign key(event_id) references events(id),
	foreign key(user_id) references users(id)
	)
	`
	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic("could not create registration table")
	}
}
