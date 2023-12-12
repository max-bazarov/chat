package app

import (
	"github.com/labstack/echo/v4"
)

func Run(port string, e *echo.Echo) error {
	return e.Start(port)
}
