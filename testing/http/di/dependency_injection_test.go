package di

import (
	"fmt"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockEmail struct {
}

func (mockEmail) Send(msg string) {
	fmt.Println("FAKE sent email: ", msg)
}

func TestNewServeMuxDI(t *testing.T) {
	is := is.New(t)

	mockEmail := mockEmail{}
	srv := NewServeMuxDI(mockEmail)

	req, err := http.NewRequest("GET", "/path", nil)
	is.NoErr(err)

	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)

	is.Equal(w.Code, http.StatusOK)
	is.Equal(w.Body.String(), `{"name": "anuchit"}`)
}
