package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

func createUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func NewEcho() *echo.Echo {
	e := echo.New()
	e.POST("/users", createUser)
	return e
}
