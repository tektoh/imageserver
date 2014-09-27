package imageserver

import (
    "testing"
    "net/http"
    "net/http/httptest"
)

type TestDispatcher struct {
    methods map[string]interface{}
}

func NewTestDispatcher() *TestDispatcher {
    dispatch := &TestDispatcher {
        methods: map[string]interface{}{
            "GET": func(req *http.Request) bool {
                return true
            },
        },
    }
    return dispatch
}

func (dispatcher *TestDispatcher) Dispatch(req *http.Request) (int, interface{}) {
    return http.StatusOK, ""
}

func TestDispatchHandler(t *testing.T) {
	ts := httptest.NewServer(Handler(NewTestDispatcher()))
    defer ts.Close()

    var res *http.Response
    var err error

    res, err = http.Get(ts.URL + "/")
    if err != nil {
        t.Error("http.Get Error", err)
        return
    }

    if res.StatusCode != http.StatusOK {
        t.Error("Get status code error", res.StatusCode)
        return
    }
}
