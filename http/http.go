package http

import (
	"github.com/labstack/echo/v4"
)

var e = echo.New()

func Start() {
	e.Logger.Fatal(e.Start(":1323"))
}
