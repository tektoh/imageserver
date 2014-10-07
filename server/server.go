package server

import (
	"net/http"
)

func Run() {
	Handle()
	http.ListenAndServe(":8080", nil)
}

func Handle() {
	http.Handle("/", NewRouter())
}
