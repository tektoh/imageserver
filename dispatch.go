package imageserver

import (
	"net/http"
)

type Dispatcher interface {
	Dispatch(req *http.Request) (int, interface{})
}

type RequestDispatcher struct {
	methods map[string]interface{}
}

func NewRequestDispatcher() *RequestDispatcher {
	dispatch := &RequestDispatcher {
		methods: map[string]interface{}{
			"GET": Get,
			"POST": Post,
			"PUT": Put,
			"DELETE": Delete,
		},
	}
	return dispatch
}

func (dispatcher *RequestDispatcher) IsAllowedMethod(method string) bool {
	for m, _ := range dispatcher.methods {
		if method == m {
			return true
		}
	}
	return false
}

func (dispatcher *RequestDispatcher) Dispatch(req *http.Request) (int, interface{}) {
	if !dispatcher.IsAllowedMethod(req.Method) {
		return http.StatusMethodNotAllowed, ""
	}
	f := dispatcher.methods[req.Method]
	return f.(func(r *http.Request) (int, interface{}))(req)
}

func Get(req *http.Request) (int, interface{}) {
	return 200, ""
}

func Post(req *http.Request) (int, interface{}) {
	return 200, ""
}

func Put(req *http.Request) (int, interface{}) {
	return 200, ""
}

func Delete(req *http.Request) (int, interface{}) {
	return 200, ""
}
