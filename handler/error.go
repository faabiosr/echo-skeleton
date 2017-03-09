package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type (
	ErrResponse struct {
		Error ErrContent `json:"error"`
	}

	ErrContent struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func Error(err error, c echo.Context) {
	code := http.StatusServiceUnavailable
	msg := http.StatusText(code)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
	}

	if c.Echo().Debug {
		msg = err.Error()
	}

	c.Logger().Error(err)

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			c.NoContent(code)
		} else {
			c.JSON(code, &ErrResponse{ErrContent{code, msg}})
		}
	}
}
