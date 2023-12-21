package rest

import "github.com/labstack/echo/v4"

func LoadRoute(e *echo.Echo, handler *handler) {
	e.GET("/menu", handler.GetMenu)
}
