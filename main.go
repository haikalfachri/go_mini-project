package main

import (
	"context"

	echo "github.com/labstack/echo/v4"
)

type operation func(ctx context.Context) error

func main() {
	e := echo.New()

	e.Logger.Fatal(e.Start(":8000"))

}

