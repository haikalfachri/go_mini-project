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
	"golang.org/x/crypto/bcrypt"
)

var userController *UserController = InitUserContoller(&configJWT)

func TestRegister_Success(t *testing.T){
	testcase := TestCase{
		name:                   "success",
		path:                   "/register",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	hashedPass, err := bcrypt.GenerateFromPassword([]byte("test123"), bcrypt.DefaultCost) 
	
	if err != nil {
		t.Errorf("error create password: %v\n", err)
	}

	var userInput input.UserInput = input.UserInput{
		Name: "test",
		Email: "test@test.com",
		Password: string(hashedPass),
		Role: "admin",
	}

	formData := url.Values{}
	formData.Set("name", userInput.Name)
	formData.Set("email", userInput.Email)
	formData.Set("password", userInput.Password)
	formData.Set("role", userInput.Role)

	bodyReader := strings.NewReader(formData.Encode())

	req := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)

	if assert.NoError(t, userController.Register(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestRegister_Failed(t *testing.T){
	testcase := TestCase{
		name:                   "failed",
		path:                   "/register",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	hashedPass, err := bcrypt.GenerateFromPassword([]byte("test123"), bcrypt.DefaultCost) 
	
	if err != nil {
		t.Errorf("error create password: %v\n", err)
	}

	var userInput input.UserInput = input.UserInput{
		Name: "test",
		Email: "test@test.com",
		Password: string(hashedPass),
		// empty role will gives error because all field must be filled
		Role: "",
	}

	formData := url.Values{}
	formData.Set("name", userInput.Name)
	formData.Set("email", userInput.Email)
	formData.Set("password", userInput.Password)
	formData.Set("role", userInput.Role)

	bodyReader := strings.NewReader(formData.Encode())

	req := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)

	if assert.NoError(t, userController.Register(ctx)) {
		assert.Equal(t, testcase.expectedStatus, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		
		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestLogin_Success(t *testing.T) {
	testcase := TestCase{
		name:                   "success",
		path:                   "/login",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	user, err := database.SeedUser(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	var userInput input.UserInput = input.UserInput{
		Name: user.Name,
		Email: user.Email,
		Password: "test123",
	}

	formData := url.Values{}
	formData.Set("email", userInput.Email)
	formData.Set("password", userInput.Password)

	bodyReader := strings.NewReader(formData.Encode())

	req := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)

	if assert.NoError(t, userController.Login(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}

func TestLogin_Failed(t *testing.T) {
	testcase := TestCase{
		name:                   "faield",
		path:                   "/login",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	user, err := database.SeedUser(database.ConnectDB())

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	var userInput input.UserInput = input.UserInput{
		Name: user.Name,
		Email: user.Email,
		// wrong password will gives an error
		Password: "wrong password",
	}

	formData := url.Values{}
	formData.Set("email", userInput.Email)
	formData.Set("password", userInput.Password)

	bodyReader := strings.NewReader(formData.Encode())

	req := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(testcase.path)

	if assert.NoError(t, userController.Login(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders(database.ConnectDB())
	})
}



