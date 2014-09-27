package main

import (
  "github.com/tektoh/imageserver"
  "net/http"
)

func main() {
  http.HandleFunc("/", imageserver.Handler)
  http.ListenAndServe(":8080", nil)
}
