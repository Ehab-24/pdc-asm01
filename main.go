package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
  port := 3000

  handleHTTP := http.NewServeMux()
  handleHTTP.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
    fmt.Fprintf(w, "Ehab: How is everything going?")
  })

  err := http.ListenAndServe(fmt.Sprintf(":%d", port), handleHTTP)
  if err != nil {
    log.Fatal(err)
  }
}
