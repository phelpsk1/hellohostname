package main

import (
  "os"
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  name, err := os.Hostname()
  if err != nil {
    fmt.Fprintf(w, "Oops: %v\n", err)
    return
  }
  fmt.Fprintf(w, "Hello! I am %s", name)
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":30080",nil)
}
