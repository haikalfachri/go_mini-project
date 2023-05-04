package controllers

import (
	"mini_project/middlewares"
	"mini_project/models/response"
	"mini_project/models/input"
	"mini_project/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type VehicleController struct {
	service services.VehicleService
}

func InitVehicleContoller(jwtAuth *middlewares.JWTConfig) *VehicleController {
	return &VehicleController{
		service: services.InitVehicleService(jwtAuth),
	}
}

func (uc *VehicleController) Create(c echo.Context) error {
	var vehicleInput input.VehicleInput
	c.Bind(&vehicleInput)

	err := vehicleInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "input invalid",
			Error	:  err.Error(),
		})
	}

	vehicle, err := uc.service.Create(vehicleInput)
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
		Data	:  vehicle,
	})
}

func (uc *VehicleController) GetByName(c echo.Context) error {
	name := c.FormValue("name")

	vehicles, err := uc.service.GetByName(name)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch a vehicle by name",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch a vehicle by name",
		Data	:  vehicles,
	})
}

func (uc *VehicleController) GetAll(c echo.Context) error {
	vehicles, err := uc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch all vehicle",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch all vehicle",
		Data	:  vehicles,
	})
}

func (uc *VehicleController) GetById(c echo.Context) error {
	id := c.Param("id")

	vehicle, err := uc.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch a vehicle by id",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch a vehicle by id",
		Data	:  vehicle,
	})
}

func (uc *VehicleController) UpdateRating(c echo.Context) error {
	id := c.Param("id")

	vehicle, err := uc.service.UpdateRating(id)
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to update rating of a vehicle",
			Error	:  err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to update rating of a vehicle",
		Data	:  vehicle,
	})
}

func (uc *VehicleController) Update(c echo.Context) error {
	var vehicleInput input.VehicleInput
	c.Bind(&vehicleInput)

	err := vehicleInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "input invalid",
			Error	:  err.Error(),
		})
	}

	id := c.Param("id")
	vehicle, err := uc.service.Update(vehicleInput, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to update a vehicle",
			Error	:  err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to update a vehicle",
		Data	:  vehicle,
	})
}

func (uc *VehicleController) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := uc.service.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to delete a vehicle",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to delete a vehicle",
	})
}

