package routes

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware   		echo.MiddlewareFunc
	JWTMiddleware      		echojwt.Config
}

func (cl *ControllerList) SetUpRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	e.Group("/auth", echojwt.WithConfig(cl.JWTMiddleware))
}