package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func init() {
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello", "World")
	})
}
