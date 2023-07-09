package database 

import (
  "database/sql"
  "fmt"
  _ "github.com/mattn/go-sqlite3"
)

func DatabaseInit(dbPath string) error {
  fmt.Println("Database initializing...")

  db, err := sql.Open("sqlite3", dbPath)

  if err != nil {
    return fmt.Errorf("Database creation failed: error %w", err)
  }

  userTableSQL := `
  CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT
  )`

  _, userTableErr := db.Exec(userTableSQL)

  if userTableErr != nil {
    fmt.Println("Failed to create table. Reason: %w", err)
  }

  defer db.Close()

  fmt.Println("Database created successfully!")
  return nil
}
