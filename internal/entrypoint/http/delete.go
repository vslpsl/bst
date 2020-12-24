package http

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

type DeletionHandler interface {
	Delete(value int) error
}

func deletePlz(handler DeletionHandler) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.QueryParam(valueQueryParam)
		intValue, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		return handler.Delete(intValue)
	}
}
