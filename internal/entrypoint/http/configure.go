package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	valueQueryParam = "val"
)

type Handler interface {
	SearchHandler
	InsertionHandler
	DeletionHandler
}

func Configure(echo *echo.Echo, handler Handler) {
	echo.Use(mapError())
	echo.Use(middleware.Logger())
	echo.GET("/search", search(handler))
	echo.POST("/insert", insert(handler))
	echo.DELETE("/delete", deletePlz(handler))
}
