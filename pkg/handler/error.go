package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type (
	//ErrResponse for dispatching errors in json format
	ErrResponse struct {
		Error ErrContent `json:"error"`
	}

	//ErrContent for dispatching error content in json format
	ErrContent struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

// Error handle the default echo erro to dispatch in json format
func Error(err error, c echo.Context) {
	code := http.StatusServiceUnavailable
	msg := http.StatusText(code)
	logMessage := err.Error()

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
		logMessage = msg
	}

	if c.Echo().Debug {
		msg = err.Error()
	}

	content := map[string]interface{}{
		"request_id": c.Response().Header().Get(echo.HeaderXRequestID),
		"message":    logMessage,
		"status":     code,
	}

	c.Logger().Errorj(content)

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			c.NoContent(code)
		} else {
			c.JSON(code, &ErrResponse{ErrContent{code, msg}})
		}
	}
}
