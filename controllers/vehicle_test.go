package controllers

import (
	"fmt"
	"mini_project/database"
	"mini_project/models/input"
	"strconv"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/stretchr/testify/assert"
)

var vehicleController *VehicleController = InitVehicleContoller(&configJWT)

func TestCreateVehicle_Success(t *testing.T){
	testcase := TestCase{
		name:                   "success",
		path:                   "/auth/vehicles",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	var vehicleInput input.VehicleInput = input.VehicleInput{
		NumberPlate: "test",
		Type: "motorcycle",
		Name: "test",
		Price: 150000.00,
		Rating: 0.00,
	}

	formData := url.Values{}
	formData.Set("number_plate", vehicleInput.NumberPlate)
	formData.Set("vehicle_type", vehicleInput.Type)
	formData.Set("name", vehicleInput.Name)
	str_price := strconv.FormatFloat(vehicleInput.Price, 'f', 2, 64)
	formData.Set("price", str_price)
	str_rating := strconv.FormatFloat(vehicleInput.Rating, 'f', 2, 64)
	formData.Set("rating", str_rating)

	bodyReader := strings.NewReader(formData.Encode())

	req := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)

	if assert.NoError(t, vehicleController.Create(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestCreateVehicle_Failed(t *testing.T){
	testcase := TestCase{
		name:                   "failed",
		path:                   "/auth/vehicles",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	var vehicleInput input.VehicleInput = input.VehicleInput{
		NumberPlate: "test",
		Type: "motorcycle",
		Name: "test",
		// price 0 gives error because default value of gorm
		Price: 0.00,
		Rating: 0.00,
	}

	formData := url.Values{}
	formData.Set("number_plate", vehicleInput.NumberPlate)
	formData.Set("vehicle_type", vehicleInput.Type)
	formData.Set("name", vehicleInput.Name)
	str_price := strconv.FormatFloat(vehicleInput.Price, 'f', 2, 64)
	formData.Set("price", str_price)
	str_rating := strconv.FormatFloat(vehicleInput.Rating, 'f', 2, 64)
	formData.Set("rating", str_rating)

	bodyReader := strings.NewReader(formData.Encode())

	req := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)

	if assert.NoError(t, vehicleController.Create(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestGetByNameVehicle_Success(t *testing.T){
	testcase := TestCase{
		name:                   "success",
		path:                   "/auth/vehicles/name",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	vehicle, err := database.SeedVehicle(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	formData := url.Values{}
	formData.Set("name", vehicle.Name)

	bodyReader := strings.NewReader(formData.Encode())

	req := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)

	if assert.NoError(t, vehicleController.GetByName(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestGetAllVehicle_Success(t *testing.T) {
	testcase := TestCase{
		name:                   "success",
		path:                   "/auth/vehicles",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	_, err := database.SeedVehicle(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)

	if assert.NoError(t, vehicleController.GetAll(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestGetByIdVehicle_Success(t *testing.T) {
	testcase := TestCase{
		name:                   "success",
		path:                   "/auth/vehicles/:vehicleId",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	vehicle, err := database.SeedVehicle(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.FormatUint(uint64(vehicle.ID), 10))

	if assert.NoError(t, vehicleController.GetById(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestGetByIdVehicle_Failed(t *testing.T) {
	testcase := TestCase{
		name:                   "failed",
		path:                   "/auth/vehicles/:vehicleId",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	_, err := database.SeedVehicle(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	// vehicle with id = 0 not exists
	ctx.SetParamValues("0")

	if assert.NoError(t, vehicleController.GetById(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestUpdateRatingVehicle_Success(t *testing.T) {
	testcase := TestCase{
		name:                   "success",
		path:                   "/auth/vehicles/rate/:vehicleId",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	order, err := database.SeedOrder(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req := httptest.NewRequest(http.MethodPut, testcase.path, nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.FormatUint(uint64(order.VehicleID), 10))

	if assert.NoError(t, vehicleController.UpdateRating(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestUpdateRatingVehicle_Failed(t *testing.T) {
	testcase := TestCase{
		name:                   "success",
		path:                   "/auth/vehicles/rate/:vehicleId",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	_, err := database.SeedOrder(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req := httptest.NewRequest(http.MethodPut, testcase.path, nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	// vehicle with id = 0 not exists
	ctx.SetParamValues("0")

	if assert.NoError(t, vehicleController.UpdateRating(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestUpdateVehicle_Success(t *testing.T) {
	testcase := TestCase{
		name:                   "success",
		path:                   "/auth/vehicles/:vehicleId",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	vehicle, err := database.SeedVehicle(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	var vehicleInput input.VehicleInput = input.VehicleInput{
		NumberPlate: "test",
		Type: "motorcycle",
		Name: "test",
		Price: 150000.00,
	}

	formData := url.Values{}
	formData.Set("number_plate", vehicleInput.NumberPlate)
	formData.Set("vehicle_type", vehicleInput.Type)
	formData.Set("name", vehicleInput.Name)
	str_price := strconv.FormatFloat(vehicleInput.Price, 'f', 2, 64)
	formData.Set("price", str_price)

	bodyReader := strings.NewReader(formData.Encode())

	req := httptest.NewRequest(http.MethodPut, testcase.path, bodyReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.FormatUint(uint64(vehicle.ID), 10))

	if assert.NoError(t, vehicleController.Update(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestUpdateVehicle_Failed(t *testing.T) {
	testcase := TestCase{
		name:                   "failed",
		path:                   "/auth/vehicles/:vehicleId",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	vehicle, err := database.SeedVehicle(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	var vehicleInput input.VehicleInput = input.VehicleInput{
		NumberPlate: "test",
		Type: "motorcycle",
		Name: "test",
		// error when update price with value 0 because of default value gorm
		Price: 0.00,
	}

	formData := url.Values{}
	formData.Set("number_plate", vehicleInput.NumberPlate)
	formData.Set("vehicle_type", vehicleInput.Type)
	formData.Set("name", vehicleInput.Name)
	str_price := strconv.FormatFloat(vehicleInput.Price, 'f', 2, 64)
	formData.Set("price", str_price)

	bodyReader := strings.NewReader(formData.Encode())

	req := httptest.NewRequest(http.MethodPut, testcase.path, bodyReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.FormatUint(uint64(vehicle.ID), 10))

	if assert.NoError(t, vehicleController.Update(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestDeleteVehicle_Success(t *testing.T) {
	testcase := TestCase{
		name:                   "success",
		path:                   "/auth/vehicles/:vehicleId",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	vehicle, err := database.SeedVehicle(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req := httptest.NewRequest(http.MethodDelete, testcase.path, nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.FormatUint(uint64(vehicle.ID), 10))

	if assert.NoError(t, vehicleController.Delete(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestDeleteVehicle_Failed(t *testing.T) {
	testcase := TestCase{
		name:                   "failed",
		path:                   "/auth/vehicles/:vehicleId",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	_, err := database.SeedVehicle(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req := httptest.NewRequest(http.MethodDelete, testcase.path, nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	// vehicle with id = 0 not exists
	ctx.SetParamValues("0")

	if assert.NoError(t, vehicleController.Delete(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}



