package main

import (
	"context"
	"log"
  	"mini_project/controllers"
	"mini_project/database"
	"mini_project/middlewares"
	"mini_project/routes"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	echo "github.com/labstack/echo/v4"
)

type operation func(ctx context.Context) error

func main() {
	db := database.ConnectDB()

	database.MigrateDB(db)

	configLogger := middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	configJWT := middlewares.JWTConfig{
		SecretKey		: "secret_key",
		ExpiresDuration	: 1,
	}

	e := echo.New()

	userCtrl := controllers.InitUserContoller(&configJWT)
	vehicleCtrl := controllers.InitVehicleContoller(&configJWT)

	routesInit := routes.ControllerList{
		LoggerMiddleware:   configLogger.Init(),
		JWTMiddleware	:   configJWT.Init(),
		UserController	: 	*userCtrl,
		VehicleController:  *vehicleCtrl,

	}

	routesInit.SetUpRoutes(e)

	go func() {
		if err := e.Start(":8000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"database": func(ctx context.Context) error {
			return database.CloseDB(db)
		},
		"http-server": func(ctx context.Context) error {
			return e.Shutdown(context.Background())
		},
	})

	<-wait
}

func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscalls that you want to be notified with
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("shutting down")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		// Do the operations asynchronously to save time
		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%s: clean up failed: %s", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}

