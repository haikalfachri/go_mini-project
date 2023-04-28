package routes

import (
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware   		echo.MiddlewareFunc
}

func (cl *ControllerList) SetUpRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)
}