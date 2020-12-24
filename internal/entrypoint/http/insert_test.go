package http

import (
	"bst/internal/app"
	"bst/pkg/model"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_insert(t *testing.T) {
	tree := model.NewTree([]int{1, 2, 3, 4, 5})
	a := app.New(tree)
	e := echo.New()
	e.Use(mapError())

	requestBody1 := &insertRequest{Val: 3}
	data, err := json.Marshal(requestBody1)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/insert", bytes.NewReader(data))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.Equal(t, model.ErrAlreadyExists, insert(a)(c))

	requestBody2 := &insertRequest{Val: 24}
	data, err = json.Marshal(requestBody2)
	assert.NoError(t, err)

	req = httptest.NewRequest(http.MethodGet, "/insert", bytes.NewReader(data))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	assert.NoError(t, insert(a)(c))
}
