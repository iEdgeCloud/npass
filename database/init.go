package database

import (
    "database/sql"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "./users.db")
    if err != nil {
        log.Fatalf("Error opening database: %v\n", err)
    }

    createTableQuery := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );`

    _, err = DB.Exec(createTableQuery)
    if err != nil {
        log.Fatalf("Error creating table: %v\n", err)
    } else {
        log.Println("Table created successfully or already exists.")
    }
}

