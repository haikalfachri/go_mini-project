package controllers

import (
	"mini_project/middlewares"
	"mini_project/models/response"
	"mini_project/models/input"
	"mini_project/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.UserService
}

func InitUserContoller(jwtAuth *middlewares.JWTConfig) *UserController {
	return &UserController{
		service: services.InitUserService(jwtAuth),
	}
}

func (uc *UserController) Register(c echo.Context) error {
	var userInput input.UserInput
	c.Bind(&userInput)

	err := userInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "input invalid",
			Error	:  err.Error(),
		})
	}

	user, err := uc.service.Register(userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "request invalid",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "register success",
		Data	:  user,
	})
}

func (uc *UserController) Login(c echo.Context) error {
	var userInput input.UserInput
	c.Bind(&userInput)

	token, err := uc.service.Login(userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "login failed",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "login success",
		Data	:  token,
	})
}


