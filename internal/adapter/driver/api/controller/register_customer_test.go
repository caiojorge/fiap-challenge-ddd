package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostRegisterCustomer(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize the router
	r := gin.Default()
	r.POST("/register", PostRegisterCustomer)

	// Create a JSON body
	requestBody := bytes.NewBuffer([]byte(`{"cpf":"123.456.789-09", "name":"John Doe","email":"johndoe@example.com"}`))

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
	assert.Equal(t, http.StatusOK, w.Code, "Expected response code to be 200")
	assert.Contains(t, w.Body.String(), "customer created John Doe", "Response body should contain correct customer name")
}
