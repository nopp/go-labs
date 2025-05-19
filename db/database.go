package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
    var err error
    DB, err = sql.Open("sqlite3", "./database.db")
    if err != nil {
        return err
    }
    return createTables()
}

func createTables() error {
    _, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );`)
    if err != nil {
        return err
    }

    _, err = DB.Exec(`CREATE TABLE IF NOT EXISTS sessions (
        id TEXT PRIMARY KEY,
        username TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`)
    return err
}

func CreateDefaultUser() error {
    _, err := DB.Exec(`INSERT OR IGNORE INTO users(username, password) VALUES (?, ?)`, "admin", "1234")
    return err
}