package x

import (
	"net/http"
	"testing"
)

func TestX(t *testing.T) {
	x := New("")
	x.GET("/boom", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("BOOM"))
		if err != nil {
			return
		}
	})
	err := x.Run()
	if err != nil {
		return
	}
}
