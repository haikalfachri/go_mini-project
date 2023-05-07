package controllers

import (
	"fmt"
	"mini_project/database"
	"mini_project/models/input"
	"mini_project/models"
	"strconv"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/stretchr/testify/assert"
)

var orderController *OrderController = InitOrderContoller(&configJWT)

func TestCreateOrder_Success(t *testing.T){
	testcase := TestCase{
		name:                   "success",
		path:                   "/auth/orders",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	user, err := database.SeedUser(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	vehicle, err := database.SeedVehicle(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	transaction := models.Transaction{
		Name: "test.img",
		Data: []byte{},
	}

	result := database.ConnectDB().Create(&transaction)

	if err := result.Error; err != nil {
		t.Errorf("error: %v\n", err)
	}

	if err := result.Last(&transaction).Error; err != nil {
		t.Errorf("error: %v\n", err)
	}

	var order models.Order = models.Order{
		UserID: user.ID,
		VehicleID: vehicle.ID,
		TransactionID: transaction.ID,
		Transaction: transaction,
		RentDuration: 1,
		Status: "pending",
		OrderRate: 0.0,
	}

	var orderInput input.OrderInput = input.OrderInput{
		UserID: order.UserID,
		VehicleID: order.VehicleID,
		RentDuration: order.RentDuration,
		Status: order.Status,
	}

	formData := url.Values{}
	str_user_id := strconv.FormatUint(uint64(orderInput.UserID), 10)
	formData.Set("user_id", str_user_id)
	str_vehicle_id := strconv.FormatUint(uint64(orderInput.VehicleID), 10)
	formData.Set("vehicle_id", str_vehicle_id)
	str_rent_duration := strconv.FormatInt(int64(orderInput.RentDuration), 10)
	formData.Set("rent_duration", str_rent_duration)
	formData.Set("status", orderInput.Status)

	bodyReader := strings.NewReader(formData.Encode())

	req := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)

	if assert.NoError(t, orderController.Create(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestCreateOrder_Failed(t *testing.T){
	testcase := TestCase{
		name:                   "failedd",
		path:                   "/auth/orders",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	user, err := database.SeedUser(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	vehicle, err := database.SeedVehicle(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	transaction := models.Transaction{
		Name: "test.img",
		Data: []byte{},
	}

	result := database.ConnectDB().Create(&transaction)

	if err := result.Error; err != nil {
		t.Errorf("error: %v\n", err)
	}

	if err := result.Last(&transaction).Error; err != nil {
		t.Errorf("error: %v\n", err)
	}

	var order models.Order = models.Order{
		UserID: user.ID,
		VehicleID: vehicle.ID,
		TransactionID: transaction.ID,
		Transaction: transaction,
		RentDuration: 1,
		Status: "pending",
		OrderRate: 0.0,
	}

	var orderInput input.OrderInput = input.OrderInput{
		UserID: order.UserID,
		VehicleID: order.VehicleID,
		// rent duration's value minimum is 1 
		RentDuration: 0,
		Status: order.Status,
	}

	formData := url.Values{}
	str_user_id := strconv.FormatUint(uint64(orderInput.UserID), 10)
	formData.Set("user_id", str_user_id)
	str_vehicle_id := strconv.FormatUint(uint64(orderInput.VehicleID), 10)
	formData.Set("vehicle_id", str_vehicle_id)
	str_rent_duration := strconv.FormatInt(int64(orderInput.RentDuration), 10)
	formData.Set("rent_duration", str_rent_duration)
	formData.Set("status", orderInput.Status)

	bodyReader := strings.NewReader(formData.Encode())

	req := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)

	if assert.NoError(t, orderController.Create(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestGetAllOrder_Success(t *testing.T){
	testcase := TestCase{
		name:                   "success",
		path:                   "/auth/orders",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	_, err := database.SeedOrder(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)

	if assert.NoError(t, orderController.GetAll(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestGetHistoryOrder_Success(t *testing.T){
	testcase := TestCase{
		name:                   "success",
		path:                   "/auth/orders/history/:userId",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	order, err := database.SeedOrder(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.FormatUint(uint64(order.UserID), 10))

	if assert.NoError(t, orderController.GetHistory(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestGetHistoryOrder_Failed(t *testing.T){
	testcase := TestCase{
		name:                   "failed",
		path:                   "/auth/orders/history/:id",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	_, err := database.SeedOrder(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	// user with id = 0 not exists
	ctx.SetParamValues("0")

	if assert.NoError(t, orderController.GetHistory(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestGiveRatingOrder_Success(t *testing.T){
	testcase := TestCase{
		name:                   "success",
		path:                   "/auth/orders/rate/:id",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	order, err := database.SeedOrder(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	order.Status = "accepted"

	if err := database.ConnectDB().Save(&order).Error; err != nil {
		t.Errorf("error: %v\n", err)
	}

	var orderInput input.OrderInput = input.OrderInput{
		OrderRate: 4.5,
	}

	formData := url.Values{}
	str_order_rate := strconv.FormatFloat(orderInput.OrderRate, 'f', 2, 64)
	formData.Set("order_rate", str_order_rate)

	bodyReader := strings.NewReader(formData.Encode())

	req := httptest.NewRequest(http.MethodPut, testcase.path, bodyReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.FormatUint(uint64(order.ID), 10))

	if assert.NoError(t, orderController.GiveRating(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestGiveRatingOrder_Failed(t *testing.T){
	testcase := TestCase{
		name:                   "failed",
		path:                   "/auth/orders/rate/:id",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	order, err := database.SeedOrder(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	// pending status gives error because status accepted is required to give rating of an order
	order.Status = "pending"

	var orderInput input.OrderInput = input.OrderInput{
		OrderRate: 4.5,
	}

	formData := url.Values{}
	str_order_rate := strconv.FormatFloat(orderInput.OrderRate, 'f', 2, 64)
	formData.Set("order_rate", str_order_rate)

	bodyReader := strings.NewReader(formData.Encode())

	req := httptest.NewRequest(http.MethodPut, testcase.path, bodyReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.FormatUint(uint64(order.ID), 10))

	if assert.NoError(t, orderController.GiveRating(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}


