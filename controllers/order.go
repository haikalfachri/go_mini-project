package controllers

import (
	"mini_project/middlewares"
	"mini_project/models/response"
	"mini_project/models/input"
	"mini_project/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	service services.OrderService
}

func InitOrderContoller(jwtAuth *middlewares.JWTConfig) *OrderController {
	return &OrderController{
		service: services.InitOrderService(jwtAuth),
	}
}

func (uc *OrderController) Create(c echo.Context) error {
	var orderInput input.OrderInput
	c.Bind(&orderInput)

	err := orderInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "input invalid",
			Error	:  err.Error(),
		})
	}

	order, err := uc.service.Create(orderInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "request invalid",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "create success",
		Data	:  order,
	})
}

func (uc *OrderController) GetByName(c echo.Context) error {
	orders, err := uc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch a order by name",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch a order by name",
		Data	:  orders,
	})
}

func (uc *OrderController) GetAll(c echo.Context) error {
	orders, err := uc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch all order",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch all order",
		Data	:  orders,
	})
}

func (uc *OrderController) GetById(c echo.Context) error {
	id := c.Param("id")

	order, err := uc.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch a order by id",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch a order by id",
		Data	:  order,
	})
}

func (uc *OrderController) GiveRating(c echo.Context) error {
	var orderInput input.OrderInput
	c.Bind(&orderInput)

	id := c.Param("id")
	order, err := uc.service.UpdateRating(orderInput, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to give rating of an order",
			Error	:  err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to give rating of an order",
		Data	:  order,
	})
}

func (uc *OrderController) Update(c echo.Context) error {
	var orderInput input.OrderInput
	c.Bind(&orderInput)

	err := orderInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "input invalid",
			Error	:  err.Error(),
		})
	}

	id := c.Param("id")
	order, err := uc.service.Update(orderInput, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to update a order",
			Error	:  err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to update a order",
		Data	:  order,
	})
}

func (uc *OrderController) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := uc.service.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to delete a order",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to delete a order",
	})
}

