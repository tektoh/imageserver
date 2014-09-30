package imageserver

import (
	"net/http"
)

func Run() {
	http.Handle("/", NewRouter())
	http.ListenAndServe(":8080", nil)
}
