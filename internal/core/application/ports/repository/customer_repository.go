package repository

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

// CustomerRepository defines the methods for interacting with the customer data.
type CustomerRepository interface {
	// CreateCustomer creates a new customer.
	CreateCustomer(ctx context.Context, customer *entity.Customer) error

	// GetCustomerByID retrieves a customer by ID.
	GetCustomerByID(ctx context.Context, id string) (*entity.Customer, error)

	// ListCustomers retrieves a list of customers.
	ListCustomers(ctx context.Context) ([]*entity.Customer, error)
}
