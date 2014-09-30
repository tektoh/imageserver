package imageserver

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func GetHandler(w http.ResponseWriter, req *http.Request) {
	args := mux.Vars(req)

	id  := args["id"]
	fnc := args["fnc"]
	ext := args["ext"]

	fmt.Fprintf(w, "GET id: %s, fnc: %s, ext: %s", id, fnc, ext)
}

func PostHandler(w http.ResponseWriter, req *http.Request) {
}

func DeleteHandler(w http.ResponseWriter, req *http.Request) {
}
