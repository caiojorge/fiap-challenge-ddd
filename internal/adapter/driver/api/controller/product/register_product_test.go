package controller

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestRegisterProductController(t *testing.T) {
	// repo := NewMockCustomerRepository()
	// mock := NewMockRegisterCustomerUseCase(repo)

	controller := NewRegisterProductController(context.Background())

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize the router
	r := gin.Default()
	r.POST("/register", controller.PostRegisterProduct)

	// Create a JSON body
	requestBody := bytes.NewBuffer([]byte(
		`{
			"name":"Lanche XPTO", 
			"description":"Pão, carne, queijo e presunto",
			"category":"Lanche",
			"price": 100.00
		}`))

	// Create the HTTP request with JSON body
	req, err := http.NewRequest("POST", "/register", requestBody)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)

	// compare the response with the expected result
	assert.Equal(t, `{"name":"Lanche XPTO","description":"Pão, carne, queijo e presunto","price":100,"category":"Lanche"}`, w.Body.String())
}
