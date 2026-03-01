package db
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
var DB *sql.DB
func InitDB(){
	var err error
	DB,err = sql.Open("sqlite3","api.db")
	if err != nil{
		panic("could not connect to database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
func CreateTables(){
	createEventsTable = `
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
	
}