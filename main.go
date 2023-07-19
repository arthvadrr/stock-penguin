package main

import (
	"fmt"
	"net/http"
	"stock-penguin/database"

	"github.com/davecgh/go-spew/spew"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var message string = fmt.Sprintf("Hello server world")

	fmt.Println(message)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 8080")

	dbPath := "database/sqldata/database.db"
	err := database.DatabaseInit(dbPath)

	if err != nil {
		fmt.Println("Database creation failed. Error: %w", err)
		return
	}

	fmt.Println("hi")
	testData, err := database.FetchSymbolData("IBM")

	if err != nil {
		fmt.Println("Could not fetch symbol data. Error: ", err)
	}

	if testData != nil {
		spew.Dump(testData)
	} else {
		fmt.Println("Something's wrong..", err)
	}

	http.ListenAndServe(":8080", nil)
}
