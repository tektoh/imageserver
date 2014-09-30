package imageserver

import (
    "testing"
)

func TestNewRouter(t *testing.T) {
    router := NewRouter()
    if router == nil {
        t.Error("NewRouter error")
        return
    }
}
