package http

import (
	"bst/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type InsertionHandler interface {
	Insert(value int) (*model.Node, error)
}

type insertRequest struct {
	Val int `json:"val"`
}

func insert(handler InsertionHandler) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := &insertRequest{}
		if err := ctx.Bind(request); err != nil {
			return err
		}
		node, err := handler.Insert(request.Val)
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusCreated, node)
	}
}
