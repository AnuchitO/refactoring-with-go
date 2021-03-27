package mux

import (
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewServeMux(t *testing.T) {
	is := is.New(t)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/path", nil)
	is.NoErr(err)

	srv := NewServeMux()

	srv.ServeHTTP(w, req)

	is.Equal(w.Code, http.StatusOK)
	is.Equal(w.Body.String(), `{"name": "anuchit"}`)
}
