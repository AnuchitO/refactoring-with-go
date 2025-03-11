package gin

import (
	"github.com/labstack/echo/v4"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
)

func TestCreateUser(t *testing.T) {
	is := is.New(t)

	g := NewGin()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	g.ServeHTTP(rec, req)

	is.Equal(http.StatusCreated, rec.Code)
	is.Equal(rec.Body.String(), userJSON)
}
