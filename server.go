package imageserver

import (
	"net/http"
)

func Run() {
	http.HandleFunc("/", Handler(NewRequestDispatcher()))
	http.ListenAndServe(":8080", nil)
}
