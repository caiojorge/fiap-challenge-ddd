package controller

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestRegisterProductController(t *testing.T) {
	repo := NewMockProductRepository()
	mock := NewMockRegisterProductUseCase(repo)

	controller := NewRegisterProductController(context.Background(), mock)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize the router
	r := gin.Default()
	r.POST("/register", controller.PostRegisterProduct)

	// Create a JSON body
	requestBody := bytes.NewBuffer([]byte(`{"id": "1", "name":"Lanche XPTO","description":"Pão, carne, queijo e presunto","category":"Lanche","price": 100}`))

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
	assert.Equal(t, `{"id":"1","name":"Lanche XPTO","description":"Pão, carne, queijo e presunto","category":"Lanche","price":100}`, w.Body.String())
}

func TestPostRegisterCustomer(t *testing.T) {
	repo := NewMockProductRepository()
	mock := NewMockRegisterProductUseCase(repo)

	controller := NewRegisterProductController(context.Background(), mock)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize the router
	r := gin.Default()
	r.POST("/register", controller.PostRegisterProduct)

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
	assert.Equal(t, http.StatusOK, w.Code)
	//assert.Contains(t, w.Body.String(), "customer created John Doe", "Response body should contain correct customer name")
}

type MockRegisterProductUseCase struct {
	repository portsrepository.ProductRepository
}

func NewMockRegisterProductUseCase(repository portsrepository.ProductRepository) *MockRegisterProductUseCase {
	return &MockRegisterProductUseCase{
		repository: repository,
	}
}

func (m *MockRegisterProductUseCase) RegisterProduct(ctx context.Context, customer entity.Product) error {
	return nil
}

type MockProductRepository struct {
	mu       sync.Mutex
	products map[string]*entity.Product
}

// NewMockCustomerRepository cria uma nova instância de um MockCustomerRepository.
func NewMockProductRepository() *MockProductRepository {
	return &MockProductRepository{
		products: make(map[string]*entity.Product),
	}
}

// Create simula a criação de um novo cliente no repositório.
func (repo *MockProductRepository) Create(ctx context.Context, product *entity.Product) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.products[product.ID]; exists {
		return errors.New("customer already exists")
	}

	repo.products[product.ID] = product
	return nil
}

// Update simula a atualização de um cliente no repositório.
func (repo *MockProductRepository) Update(ctx context.Context, product *entity.Product) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.products[product.ID]; !exists {
		return errors.New("customer not found")
	}

	repo.products[product.ID] = product
	return nil
}

// Find simula a recuperação de um cliente pelo ID.
func (repo *MockProductRepository) Find(ctx context.Context, id string) (*entity.Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	customer, exists := repo.products[id]
	if !exists {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}

// FindAll simula a recuperação de uma lista de clientes.
func (repo *MockProductRepository) FindAll(ctx context.Context) ([]*entity.Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var products []*entity.Product
	for _, product := range repo.products {
		products = append(products, product)
	}
	return products, nil
}

// delete
func (repo *MockProductRepository) Delete(ctx context.Context, id string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.products[id]; !exists {
		return errors.New("product not found")
	}

	delete(repo.products, id)
	return nil
}
