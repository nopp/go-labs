package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func initDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./products.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT,
		price REAL
	);`

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}
