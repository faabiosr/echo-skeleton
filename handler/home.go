package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type (
	HomeResponse struct {
		Message string `json:"message"`
	}
)

func HomeIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, &HomeResponse{Message: "App is running!"})
}
