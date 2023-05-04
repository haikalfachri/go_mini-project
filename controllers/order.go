package controllers

import (
	"mini_project/middlewares"
	"mini_project/models/input"
	"mini_project/models/response"
	"mini_project/services"
	"net/http"
	"strconv"

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
	user_id := c.FormValue("user_id")
	uint_user_id, _ := strconv.ParseUint(user_id, 10, 64)
	vehicle_id := c.FormValue("vehicle_id")
	uint_vehicle_id, _ := strconv.ParseUint(vehicle_id, 10, 64)
	rent_duration := c.FormValue("rent_duration")
	int_vehicle_id, _ := strconv.ParseInt(rent_duration, 10, 64)
	status := c.FormValue("status")

	var orderInput input.OrderInput = input.OrderInput{
		UserID: uint(uint_user_id),
		VehicleID: uint(uint_vehicle_id),
		RentDuration: int(int_vehicle_id),
		Status: status,
	}

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

func (uc *OrderController) GetHistory(c echo.Context) error {
	id := c.Param("id")
	orders, err := uc.service.GetHistory(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch order history",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch order history",
		Data	:  orders,
	})
}

func (uc *OrderController) GiveRating(c echo.Context) error {
	order_rate := c.FormValue("order_rate")
	f64_order_rate, _ := strconv.ParseFloat(order_rate, 64)

	var orderInput input.OrderInput = input.OrderInput{
		OrderRate: f64_order_rate,
	}

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

