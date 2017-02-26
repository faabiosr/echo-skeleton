package handler

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
