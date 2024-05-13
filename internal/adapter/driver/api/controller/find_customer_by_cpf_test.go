package controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestGetCustomerByCPF tests the GetCustomerByCPF handler for both valid and invalid requests.
func TestGetCustomerByCPF(t *testing.T) {
	repo := NewMockCustomerRepository()
	mock := NewMockRegisterCustomerUseCase(repo)

	controller := NewFindCustomerByCPFController(context.Background(), mock)

	// Set up the Gin router
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Register the handler
	r.GET("/customer", controller.GetCustomerByCPF)

	// Create a test table
	tests := []struct {
		description   string
		cpfQueryParam string
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "valid cpf query",
			cpfQueryParam: "12345678901",
			expectedCode:  http.StatusOK,
			expectedBody:  `{"message":"Profile","cpf":"12345678901"}`,
		},
		{
			description:   "missing cpf query",
			cpfQueryParam: "",
			expectedCode:  http.StatusBadRequest,
			expectedBody:  `{"error":"missing cpf"}`,
		},
	}

	// Run sub-tests
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			// Create a request to pass to our handler. We don't have any query parameters for now.
			req, _ := http.NewRequest(http.MethodGet, "/customer", nil)

			// Add query parameters if needed
			if tc.cpfQueryParam != "" {
				q := req.URL.Query()
				q.Add("cpf", tc.cpfQueryParam)
				req.URL.RawQuery = q.Encode()
			}

			// Record the response
			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			// Check the status code and body
			assert.Equal(t, tc.expectedCode, resp.Code)
			assert.JSONEq(t, tc.expectedBody, resp.Body.String())
		})
	}
}

type MockFindByCPFCustomerUseCase struct {
	repository portsrepository.CustomerRepository
}

func NewMockFindByCPFCustomerUseCase(repository portsrepository.CustomerRepository) *MockFindByCPFCustomerUseCase {
	return &MockFindByCPFCustomerUseCase{
		repository: repository,
	}
}

func (m *MockRegisterCustomerUseCase) FindCustomerByCPF(ctx context.Context, cpf string) (*entity.Customer, error) {
	return &entity.Customer{}, nil
}
