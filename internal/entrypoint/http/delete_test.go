package http

import (
	"bst/internal/app"
	"bst/pkg/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_delete(t *testing.T) {
	tree := model.NewTree([]int{1, 2, 3, 4, 5})
	a := app.New(tree)

	e := echo.New()
	Configure(e, a)
	req := httptest.NewRequest(http.MethodGet, "/delete?val=1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, deletePlz(a)(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
