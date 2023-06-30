package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  var message string = fmt.Sprintf("Hello server world");

  fmt.Println(message)
}

func main() {
  http.HandleFunc("/", handler)
  fmt.Println("Server listening on port 8080")
  http.ListenAndServe(":8080", nil)
}

