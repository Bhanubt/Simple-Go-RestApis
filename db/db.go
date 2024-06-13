package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

//Connecting and Creating  to Database
func Connect() {
	var err error

	//Creating the Database
	DB, err = sql.Open("sqlite", "pvms.db")

	if err != nil {
		panic("Cannot Create Database")
	}
	
//setting the connections
	
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createAgenciesTable()
}

//Creating Agencies Table
func createAgenciesTable() {
	createAgencies := `
	CREATE TABLE IF NOT EXISTS agencies(
		agency_id TEXT NOT NULL PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		state TEXT NOT NULL
	)
	`
	//Executing the Agencies Table
	_, err := DB.Exec(createAgencies)
	if err != nil {
		panic("Error Creating Table Agencies")
	}

}
