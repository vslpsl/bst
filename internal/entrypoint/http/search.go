package http

import (
	"bst/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type SearchHandler interface {
	Search(value int) (*model.Node, error)
}

func search(handler SearchHandler) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.QueryParam(valueQueryParam)
		intValue, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		node, err := handler.Search(intValue)
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, node)
	}
}
