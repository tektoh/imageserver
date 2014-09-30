package imageserver

import (
    "testing"
    "net/http"
    "net/http/httptest"
)

func TestGetHandler(t *testing.T) {
	ts := httptest.NewServer(NewRouter())
    defer ts.Close()

    var res *http.Response
    var err error

    res, err = http.Get(ts.URL + "/0123456789abcdef/size.jpg")
    if err != nil {
        t.Error("http.Get Error", err)
        return
    }

    if res.StatusCode != http.StatusOK {
        t.Error("Get status code error", res.StatusCode)
        return
    }
}
