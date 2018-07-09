package handler

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type (
	ErrorTestSuite struct {
		HandlerTestSuite
	}
)

func (s *ErrorTestSuite) TestErrorWithGETMethod() {
	req, _ := http.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)

	s.assert.NotNil(s.echo.Router())
	Error(errors.New("error"), c)

	s.assert.Equal(http.StatusServiceUnavailable, rec.Code)
}

func (s *ErrorTestSuite) TestErrorWithHEADMethod() {
	req, _ := http.NewRequest(echo.HEAD, "/", nil)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)

	s.assert.NotNil(s.echo.Router())
	Error(errors.New("error"), c)

	s.assert.Equal(http.StatusServiceUnavailable, rec.Code)
}

func (s *ErrorTestSuite) TestErrorWithHTTPError() {
	req, _ := http.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)

	s.assert.NotNil(s.echo.Router())
	Error(echo.NewHTTPError(http.StatusBadRequest, "something"), c)

	s.assert.Equal(http.StatusBadRequest, rec.Code)
}

func (s *ErrorTestSuite) TestErrorWithDebugEnabled() {
	s.echo.Debug = true
	req, _ := http.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)

	s.assert.NotNil(s.echo.Router())
	Error(echo.NewHTTPError(http.StatusBadRequest, "something"), c)

	s.assert.Equal(http.StatusBadRequest, rec.Code)
	s.echo.Debug = false
}

func TestErrorTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorTestSuite))
}
