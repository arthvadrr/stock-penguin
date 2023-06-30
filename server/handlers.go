package main

func handler(w http.ResponseWriter, r *http.Request) {
  var message string = fmt.Sprintf("Hello server world");

  fmt.Println(message)
}
