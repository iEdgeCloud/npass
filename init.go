//package main
//
//import (
//    "log"
//    "github.com/iEdgeCloud/npass/database"
//)
//
//func main() {
//    database.InitDB()
//
//    insertUserQuery := `INSERT INTO users (username, password) VALUES (?, ?)`
//    _, err := database.DB.Exec(insertUserQuery, "user1", "pass1")
//    if err != nil {
//        log.Fatal(err)
//    }
//
//    log.Println("User user1 created")
//}
//
package main

import (
    "log"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "./users.db")
    if err != nil {
        log.Fatalf("Error opening database: %v\n", err)
    }
    defer db.Close()

    createTableQuery := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );`

    _, err = db.Exec(createTableQuery)
    if err != nil {
        log.Fatalf("Error creating table: %v\n", err)
    } else {
        log.Println("Table created successfully or already exists.")
    }

    insertUserQuery := `INSERT INTO users (username, password) VALUES (?, ?)`
    _, err = db.Exec(insertUserQuery, "user1", "pass1")
    if err != nil {
        log.Fatalf("Error inserting user: %v\n", err)
    } else {
        log.Println("User user1 created successfully.")
    }
}

