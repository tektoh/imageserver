package imageserver

import (
	"net/http"
)

func Handler(dispatcher Dispatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		code, _ := dispatcher.Dispatch(req)
		w.WriteHeader(code)
	}
}
