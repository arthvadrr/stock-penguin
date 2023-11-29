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
		fmt.Println("Failed to create table. Reason: %w", userTableErr)
	}

	symbolOverviewTableSQL := `
  CREATE TABLE IF NOT EXISTS symbols (
  symbol VARCHAR(5) PRIMARY KEY,
  assettype VARCHAR(64),
  description TEXT,
  cik INT(10),
  country CHAR(3),
  sector VARCHAR(128),
  industry TEXT
  )`

	_, symbolOverviewTableErr := db.Exec(symbolOverviewTableSQL)

	if symbolOverviewTableErr != nil {
		fmt.Println("Failed to create symbol table. Reason: %w", symbolOverviewTableErr)
	} else {
		defer db.Close()
		fmt.Println("Database created successfully!")
	}
	return nil
}
