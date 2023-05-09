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
	OrderController 		controllers.OrderController
}

func (cl *ControllerList) SetUpRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	noAuth := e.Group("")
	useAuth := e.Group("/auth", echojwt.WithConfig(cl.JWTMiddleware))
	
	noAuth.POST("/register", cl.UserController.Register)
	noAuth.POST("/login", cl.UserController.Login)

	useAuth.POST("/vehicles", cl.VehicleController.Create, middlewares.VerifyAdmin)
	useAuth.GET("/vehicles/name", cl.VehicleController.GetByName, middlewares.VerifyToken)
	useAuth.GET("/vehicles", cl.VehicleController.GetAll, middlewares.VerifyAdmin)
	// use vehicle_id
	useAuth.GET("/vehicles/:id", cl.VehicleController.GetById, middlewares.VerifyAdmin)
	useAuth.PUT("/vehicles/rate/:id", cl.VehicleController.UpdateRating, middlewares.VerifyAdmin)
	useAuth.PUT("/vehicles/:id", cl.VehicleController.Update, middlewares.VerifyAdmin)
	useAuth.DELETE("/vehicles/:id", cl.VehicleController.Delete, middlewares.VerifyAdmin)

	// use transaction_id
	useAuth.PUT("/transactions/pay/:id", cl.TransactionController.PayOrder, middlewares.VerifyToken)

	useAuth.POST("/orders", cl.OrderController.Create, middlewares.VerifyToken)
	useAuth.GET("/orders", cl.OrderController.GetAll, middlewares.VerifyAdmin)
	// use order_id
	useAuth.PUT("/orders/rate/:id", cl.OrderController.GiveRating, middlewares.VerifyToken)
	// use user_id
	useAuth.GET("/orders/history/:id", cl.OrderController.GetHistory, middlewares.VerifyToken)
}