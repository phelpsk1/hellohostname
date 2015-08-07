package main

import (
  "os"
  "fmt"
  "net/http"
	"log"
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
  port := os.Getenv("PORT")
	log.Printf("listening on %v...\n", port)
	http.HandleFunc("/", handler)
  err := http.ListenAndServe(":"+port,nil)
  if err != nil {
		panic(err)
  }
}
