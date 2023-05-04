package controllers

import (
	"bytes"
	"image"
	"image/jpeg"
	"mini_project/middlewares"
	"mini_project/models/input"
	"mini_project/models/response"
	"mini_project/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	serviceTransaction 	services.TransactionService
	serviceOrder 		services.OrderService
}

func InitTransactionContoller(jwtAuth *middlewares.JWTConfig) *TransactionController {
	return &TransactionController{
		serviceTransaction: services.InitTransactionService(jwtAuth),
		serviceOrder: services.InitOrderService(jwtAuth),
	}
}

func (uc *TransactionController) PayOrder(c echo.Context) error {
	file, err := c.FormFile("image")

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "request invalid",
			Error	:  err.Error(),
		})
	}

	src, err := file.Open()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to read the image",
			Error	:  err.Error(),
		})
	}

	defer src.Close()

	img, _, err := image.Decode(src) 

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to decode the image",
			Error	:  err.Error(),
		})
	}

	buffer := new(bytes.Buffer)

    if err := jpeg.Encode(buffer, img, nil); err != nil {
        return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to encode the image",
			Error	:  err.Error(),
		})
    }

	var transactionInput input.TransactionInput = input.TransactionInput{Name: file.Filename, Data: buffer.Bytes()}

	err = transactionInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "input invalid",
			Error	:  err.Error(),
		})
	}

	id := c.Param("id")
	transaction, err := uc.serviceTransaction.Update(transactionInput, id)
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "request invalid",
			Error	:  err.Error(),
		})
	}

	_, _ = uc.serviceOrder.UpdateStatus(id)

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "pay success",
		Data	:  transaction,
	})
}
