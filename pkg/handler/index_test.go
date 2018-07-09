package handler

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"testing"
)

type (
	IndexTestSuite struct {
		HandlerTestSuite
	}
)

func (s *IndexTestSuite) TestIndex() {
	req := httptest.NewRequest(echo.GET, "/", nil)

	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	err := Index()(ctx)
	body := `It's works`

	s.assert.NoError(err)
	s.assert.Equal(body, rec.Body.String())
}

func TestIndexTestSuite(t *testing.T) {
	suite.Run(t, new(IndexTestSuite))
}
