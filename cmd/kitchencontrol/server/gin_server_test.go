package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// mockControllerGetCustomerByCPF is a mock controller function
func mockControllerGetCustomerByCPF(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"cpf": "12345678901"})
}

// Setup the server and define URLs with mock controller
func setupTestServer() *GinServer {
	gin.SetMode(gin.TestMode)
	server := NewServer()
	server.router.GET("/customer", mockControllerGetCustomerByCPF)
	return server
}

// TestServerRoutes tests whether the routing is properly configured
func TestServerRoutes(t *testing.T) {
	server := setupTestServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/customer", nil)
	server.GetRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "12345678901")
}
