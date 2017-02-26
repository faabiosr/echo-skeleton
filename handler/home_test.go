package handler

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type (
	HomeTestSuite struct {
		HandlerTestSuite
	}
)

func (s *HomeTestSuite) TestIndex() {
	req := new(http.Request)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)
	responseJSON := `{"message":"App is running!"}`

	s.assert.NoError(HomeIndex(ctx))
	s.assert.Equal(http.StatusOK, rec.Code)
	s.assert.Equal(responseJSON, rec.Body.String())
}

func TestHomeTestSuite(t *testing.T) {
	suite.Run(t, new(HomeTestSuite))
}
