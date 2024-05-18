package usecase

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)

func TestProductRegisterAndUpdater(t *testing.T) {

	product, err := entity.NewProduct("Lanche XPTO", "Pão queijo e carne", "Lanche", 30.00)
	assert.Nil(t, err)
	assert.NotNil(t, product)

	assert.Equal(t, "Lanche XPTO", product.Name)

	assert.NotEmpty(t, product.GetID())

	repo := NewMockProductRepository()
	assert.NotNil(t, repo)

	register := NewProductRegister(repo)
	assert.NotNil(t, register)

	err = register.RegisterProduct(context.Background(), *product)
	assert.Nil(t, err)

	product2, err := repo.Find(context.Background(), product.ID)
	assert.Nil(t, err)
	assert.NotNil(t, product2)
	assert.Equal(t, product, product2)

	products, err := repo.FindAll(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Len(t, products, 1)

	updater := NewProductUpdate(repo)
	assert.NotNil(t, updater)

	product.Name = "Lanche XPTO 2"
	err = updater.UpdateProduct(context.Background(), *product)
	assert.Nil(t, err)

	product2, err = repo.Find(context.Background(), product.ID)
	assert.Nil(t, err)
	assert.NotNil(t, product2)
	assert.Equal(t, product, product2)

}

type MockProductRepository struct {
	mu       sync.Mutex
	products map[string]*entity.Product
}

// NewMockproductRepository cria uma nova instância de um MockproductRepository.
func NewMockProductRepository() *MockProductRepository {
	return &MockProductRepository{
		products: make(map[string]*entity.Product),
	}
}

// Create simula a criação de um novo cliente no repositório.
func (repo *MockProductRepository) Create(ctx context.Context, product *entity.Product) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.products[product.GetID()]; exists {
		return errors.New("product already exists")
	}

	repo.products[product.GetID()] = product
	return nil
}

// Update simula a atualização de um cliente no repositório.
func (repo *MockProductRepository) Update(ctx context.Context, product *entity.Product) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.products[product.GetID()]; !exists {
		return errors.New("product not found")
	}

	repo.products[product.GetID()] = product
	return nil
}

// Find simula a recuperação de um cliente pelo ID.
func (repo *MockProductRepository) Find(ctx context.Context, id string) (*entity.Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	product, exists := repo.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}
	return product, nil
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

func (repo *MockProductRepository) Delete(ctx context.Context, id string) (*entity.Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	product, exists := repo.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}

	delete(repo.products, id)
	return product, nil
}
