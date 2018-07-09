package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

// Index handler endpoint
func Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "It's works")
	}
}
