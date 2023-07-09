package main

import (
  "fmt"
  "net/http"
  "github.com/arthvadrr/stock-penguin/database"
)

func handler(w http.ResponseWriter, r *http.Request) {
  var message string = fmt.Sprintf("Hello server world");

  fmt.Println(message)
}

func main() {
  http.HandleFunc("/", handler)
  fmt.Println("Server listening on port 8080")
  database.DatabaseInit()
  http.ListenAndServe(":8080", nil)
}


