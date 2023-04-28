package main

import (
	"context"
	"mini_project/database"
	"mini_project/middlewares"
	"mini_project/routes"

	echo "github.com/labstack/echo/v4"
)

type operation func(ctx context.Context) error

func main() {
	db := database.ConnectDB()

	database.MigrateDB(db)

	configLogger := middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	e := echo.New()

	routesInit := routes.ControllerList{
		LoggerMiddleware:   configLogger.Init(),
	}

	routesInit.SetUpRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))

}

