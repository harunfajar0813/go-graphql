package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func Welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	}
}
