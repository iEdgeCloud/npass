package models

import (
    "database/sql"
    "errors"
    "github.com/iedgecloud/npass/database"
)

type User struct {
    ID       int
    Username string
    Password string
}

func Authenticate(username, password string) (*User, error) {
    var user User
    query := "SELECT id, username, password FROM users WHERE username = ? AND password = ?"
    err := database.DB.QueryRow(query, username, password).Scan(&user.ID, &user.Username, &user.Password)
    if err == sql.ErrNoRows {
        return nil, errors.New("invalid credentials")
    } else if err != nil {
        return nil, err
    }
    return &user, nil
}

