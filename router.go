package imageserver

import (
	"github.com/gorilla/mux"
	"net/http"
)

type handler func(w http.ResponseWriter, req *http.Request)

type Route struct {
	Method string
	Path string
	Handler handler
}

func NewRouter() http.Handler {
	var routes = []Route {
		{
			Method: "GET",
			Path: "/{id:[0-9a-f_]+}/{fnc:[a-z]+}.{ext:(jpe?g|png)}",
			Handler: GetHandler,
		},
	}

	router := mux.NewRouter()

	for _, r := range routes {
		m := router.Methods(r.Method)
		m.Path(r.Path)
		m.HandlerFunc(r.Handler)
	}
	
	return router
}
