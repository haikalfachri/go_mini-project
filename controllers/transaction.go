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

func (uc *TransactionController) Create(c echo.Context) error {
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

	transaction, err := uc.serviceTransaction.Create(transactionInput)
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
		Data	:  transaction,
	})
}

func (uc *TransactionController) GetByName(c echo.Context) error {
	transactions, err := uc.serviceTransaction.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch a transaction by name",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch a transaction by name",
		Data	:  transactions,
	})
}

func (uc *TransactionController) GetAll(c echo.Context) error {
	transactions, err := uc.serviceTransaction.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch all transaction",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch all transaction",
		Data	:  transactions,
	})
}

func (uc *TransactionController) GetById(c echo.Context) error {
	id := c.Param("id")

	transaction, err := uc.serviceTransaction.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch a transaction by id",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to fetch a transaction by id",
		Data	:  transaction,
	})
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

func (uc *TransactionController) Update(c echo.Context) error {
	var transactionInput input.TransactionInput
	c.Bind(&transactionInput)

	err := transactionInput.Validate()
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
			Message	: "failed to update a transaction",
			Error	:  err.Error(),
		})
	}


	
	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to update a transaction",
		Data	:  transaction,
	})
}

func (uc *TransactionController) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := uc.serviceTransaction.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response[any]{
			Status 	: "failed",
			Message	: "failed to delete a transaction",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[any]{
		Status 	: "success",
		Message	: "success to delete a transaction",
	})
}

