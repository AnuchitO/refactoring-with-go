package simple

import (
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSimpleHandler(t *testing.T) {
	is := is.New(t)

	req, err := http.NewRequest("GET", "/path", nil)
	is.NoErr(err)

	w := httptest.NewRecorder()
	SimpleHandler(w, req)

	is.Equal(w.Code, http.StatusOK)
	is.Equal(w.Body.String(), `{"name": "anuchit"}`)
}
