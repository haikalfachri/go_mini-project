package controllers

import (
	"mini_project/database"
	"mini_project/middlewares"

	"github.com/labstack/echo/v4"
)

type TestCase struct {
	name                   string
	path                   string
	expectedStatus         int
	expectedBodyStartsWith string
}

var configJWT middlewares.JWTConfig = middlewares.JWTConfig{
	SecretKey		: "secret_key",
	ExpiresDuration	: 1,
}

func InitEcho() *echo.Echo {
	db := database.ConnectDB()
	database.MigrateDB(db)
	e := echo.New()

	return e 
}
