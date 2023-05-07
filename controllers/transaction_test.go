package controllers

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"mini_project/database"
	"net/http"
	"net/http/httptest"
	"os"

	"strconv"
	"strings"
	"testing"

)

var transactionController TransactionController = *InitTransactionContoller(&configJWT)

func TestPayOrderTransaction_Success(t *testing.T) {
    testcase := TestCase{
        name:                   "success",
        path:                   "/auth/transactions/pay/:transactionId",
        expectedStatus:         http.StatusOK,
        expectedBodyStartsWith: "{\"status\":",
    }

    e := InitEcho()

    // create test multipart file
    requestBody := &bytes.Buffer{}

    writer := multipart.NewWriter(requestBody)

    fileWriter, err := writer.CreateFormFile("image", "test.jpg")
    if err != nil {
        t.Fatalf("failed to create form file: %v", err)
    }

    // create test image
    img := image.NewRGBA(image.Rect(0, 0, 2, 2))

	file, err := os.Create("test.jpg")
    if err != nil {
        t.Fatalf("failed to create form file: %v", err)
    }

    src, err := os.OpenFile(file.Name(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)

    defer file.Close()
    defer src.Close()

	err = jpeg.Encode(src, img, &jpeg.Options{Quality: 1})
    if err != nil {
        t.Fatalf("failed to encode file: %v", err)
    }

    _, err = io.Copy(fileWriter, file)
    if err != nil {
        t.Fatalf("failed to copy file to form file: %v", err)
    }
    
    writer.Close()

    // create test order
    order, err := database.SeedOrder(database.ConnectDB())
    if err != nil {
        t.Fatalf("error: %v", err)
    }

    req := httptest.NewRequest(http.MethodPut, "/auth/transactions/pay/:transactionId", requestBody)
    req.Header.Set("Content-Type", writer.FormDataContentType())
    req.Header.Set("Content-Length", strconv.FormatInt(int64(requestBody.Len()), 10))
    
    rec := httptest.NewRecorder()

    ctx := e.NewContext(req, rec)
    ctx.SetParamNames("id")
    ctx.SetParamValues(strconv.FormatUint(uint64(order.TransactionID), 10))

    if err := transactionController.PayOrder(ctx); err != nil {
        t.Fatalf("failed to perform request: %v", err)
    }

    if rec.Code != testcase.expectedStatus {
        t.Errorf("unexpected status code: got %v want %v", rec.Code, testcase.expectedStatus)
    }

    body := rec.Body.String()
    fmt.Println(body)

    if !strings.HasPrefix(rec.Body.String(), testcase.expectedBodyStartsWith) {
        t.Errorf("unexpected response body: got %v want %v", rec.Body.String(), testcase.expectedBodyStartsWith)
    }

    t.Cleanup(func() {
        database.CleanSeeders(database.ConnectDB())
    })
}

func TestPayOrderTransaction_Form_Failed(t *testing.T) {
    testcase := TestCase{
        name:                   "failed",
        path:                   "/auth/transactions/pay/:transactionId",
        expectedStatus:         http.StatusBadRequest,
        expectedBodyStartsWith: "{\"status\":",
    }

    e := InitEcho()

    // create test multipart file
    requestBody := &bytes.Buffer{}

    writer := multipart.NewWriter(requestBody)

    // wrong key gives invalid request error
    fileWriter, err := writer.CreateFormFile("images", "test.jpg")
    if err != nil {
        t.Fatalf("failed to create form file: %v", err)
    }

    // create test image
    img := image.NewRGBA(image.Rect(0, 0, 2, 2))

	file, err := os.Create("test.jpg")
    if err != nil {
        t.Fatalf("failed to create form file: %v", err)
    }

    // remove write permission gives open file error
    src, err := os.OpenFile(file.Name(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)

    defer file.Close()
    defer src.Close()

	err = jpeg.Encode(src, img, &jpeg.Options{Quality: 1})
    if err != nil {
        t.Fatalf("failed to encode file: %v", err)
    }

    _, err = io.Copy(fileWriter, file)
    if err != nil {
        t.Fatalf("failed to copy file to form file: %v", err)
    }
    
    writer.Close()

    // create test order
    order, err := database.SeedOrder(database.ConnectDB())
    if err != nil {
        t.Fatalf("error: %v", err)
    }

    req := httptest.NewRequest(http.MethodPut, "/auth/transactions/pay/:transactionId", requestBody)
    req.Header.Set("Content-Type", writer.FormDataContentType())
    req.Header.Set("Content-Length", strconv.FormatInt(int64(requestBody.Len()), 10))
    
    rec := httptest.NewRecorder()

    ctx := e.NewContext(req, rec)
    ctx.SetParamNames("id")
    ctx.SetParamValues(strconv.FormatUint(uint64(order.TransactionID), 10))

    if err := transactionController.PayOrder(ctx); err != nil {
        t.Fatalf("failed to perform request: %v", err)
    }

    if rec.Code != testcase.expectedStatus {
        t.Errorf("unexpected status code: got %v want %v", rec.Code, testcase.expectedStatus)
    }

    body := rec.Body.String()
    fmt.Println(body)

    if !strings.HasPrefix(rec.Body.String(), testcase.expectedBodyStartsWith) {
        t.Errorf("unexpected response body: got %v want %v", rec.Body.String(), testcase.expectedBodyStartsWith)
    }

    t.Cleanup(func() {
        database.CleanSeeders(database.ConnectDB())
    })
}

func TestPayOrderTransaction_DB_Failed(t *testing.T) {
    testcase := TestCase{
        name:                   "failed",
        path:                   "/auth/transactions/pay/:transactionId",
        expectedStatus:         http.StatusBadRequest,
        expectedBodyStartsWith: "{\"status\":",
    }

    e := InitEcho()

    // create test multipart file
    requestBody := &bytes.Buffer{}

    writer := multipart.NewWriter(requestBody)

    fileWriter, err := writer.CreateFormFile("image", "test.jpg")
    if err != nil {
        t.Fatalf("failed to create form file: %v", err)
    }

    // create test image
    img := image.NewRGBA(image.Rect(0, 0, 2, 2))

	file, err := os.Create("test.jpg")
    if err != nil {
        t.Fatalf("failed to create form file: %v", err)
    }

    src, err := os.OpenFile(file.Name(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)

    defer file.Close()
    defer src.Close()

	err = jpeg.Encode(src, img, &jpeg.Options{Quality: 1})
    if err != nil {
        t.Fatalf("failed to encode file: %v", err)
    }

    _, err = io.Copy(fileWriter, file)
    if err != nil {
        t.Fatalf("failed to copy file to form file: %v", err)
    }
    
    writer.Close()

    // create test order
    _, err = database.SeedOrder(database.ConnectDB())
    if err != nil {
        t.Fatalf("error: %v", err)
    }

    req := httptest.NewRequest(http.MethodPut, "/auth/transactions/pay/:transactionId", requestBody)
    req.Header.Set("Content-Type", writer.FormDataContentType())
    req.Header.Set("Content-Length", strconv.FormatInt(int64(requestBody.Len()), 10))
    
    rec := httptest.NewRecorder()

    ctx := e.NewContext(req, rec)
    ctx.SetParamNames("id")
    // transaction id = 0 is not exists gives record not found error
    ctx.SetParamValues("0")

    if err := transactionController.PayOrder(ctx); err != nil {
        t.Fatalf("failed to perform request: %v", err)
    }

    if rec.Code != testcase.expectedStatus {
        t.Errorf("unexpected status code: got %v want %v", rec.Code, testcase.expectedStatus)
    }

    body := rec.Body.String()
    fmt.Println(body)

    if !strings.HasPrefix(rec.Body.String(), testcase.expectedBodyStartsWith) {
        t.Errorf("unexpected response body: got %v want %v", rec.Body.String(), testcase.expectedBodyStartsWith)
    }

    t.Cleanup(func() {
        database.CleanSeeders(database.ConnectDB())
    })
}