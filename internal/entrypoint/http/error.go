package http

import (
	"bst/pkg/model"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func mapError() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if err := next(ctx); err != nil {
				if errors.Is(err, model.ErrNotFound) {
					return ctx.JSON(http.StatusNotFound, nil)
				}
				if errors.Is(err, model.ErrAlreadyExists) {
					return ctx.JSON(http.StatusConflict, nil)
				}
				return ctx.JSON(http.StatusInternalServerError, "woops")
			}
			return nil
		}
	}
}
