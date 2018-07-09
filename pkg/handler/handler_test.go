package handler

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"testing"
)

type (
	HandlerTestSuite struct {
		suite.Suite
		assert *assert.Assertions
		echo   *echo.Echo
	}
)

func (s *HandlerTestSuite) SetupTest() {
	s.assert = assert.New(s.T())
	s.echo = echo.New()
	s.echo.Logger.SetOutput(ioutil.Discard)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
