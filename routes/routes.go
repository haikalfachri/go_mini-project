package routes

import (
	"mini_project/controllers"
	"mini_project/middlewares"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware   		echo.MiddlewareFunc
	JWTMiddleware      		echojwt.Config
	UserController			controllers.UserController
	VehicleController		controllers.VehicleController
	TransactionController	controllers.TransactionController
}

func (cl *ControllerList) SetUpRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	noAuth := e.Group("")
	useAuth := e.Group("/auth", echojwt.WithConfig(cl.JWTMiddleware))
	
	noAuth.POST("/register", cl.UserController.Register)
	noAuth.POST("/login", cl.UserController.Login)

	useAuth.GET("/users", cl.UserController.GetAll, middlewares.VerifyAdmin)
	useAuth.GET("/users/:id", cl.UserController.GetById, middlewares.VerifyAdmin)
	useAuth.PUT("/users/:id", cl.UserController.Update, middlewares.VerifyAdmin)
	useAuth.DELETE("/users/:id", cl.UserController.Delete, middlewares.VerifyAdmin)

	useAuth.POST("/vehicles", cl.VehicleController.Create, middlewares.VerifyAdmin)
	useAuth.GET("/vehicles/name/:name", cl.VehicleController.GetByName, middlewares.VerifyToken)
	useAuth.GET("/vehicles", cl.VehicleController.GetAll, middlewares.VerifyAdmin)
	useAuth.GET("/vehicles/:id", cl.VehicleController.GetById, middlewares.VerifyAdmin)
	useAuth.PUT("/vehicles/:id", cl.VehicleController.Update, middlewares.VerifyAdmin)
	useAuth.DELETE("/vehicles/:id", cl.VehicleController.Delete, middlewares.VerifyAdmin)

	useAuth.POST("/transactions", cl.TransactionController.Create, middlewares.VerifyToken)
	useAuth.GET("/transactions", cl.TransactionController.GetAll, middlewares.VerifyAdmin)
	useAuth.GET("/transactions/:id", cl.TransactionController.GetById, middlewares.VerifyAdmin)
	useAuth.PUT("/transactions/:id", cl.TransactionController.Update, middlewares.VerifyAdmin)
	useAuth.DELETE("/transactions/:id", cl.TransactionController.Delete, middlewares.VerifyAdmin)
}