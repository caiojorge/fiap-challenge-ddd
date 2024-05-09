package usecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type CustomerIdentifierUseCase struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerIdentifierUseCase(customerRepo repository.CustomerRepository) *CustomerIdentifierUseCase {
	return &CustomerIdentifierUseCase{
		customerRepo: customerRepo,
	}
}

func (uc *CustomerIdentifierUseCase) IdentifyCustomer(ctx context.Context, customerID string) (*entity.Customer, error) {
	// Implement your logic here to identify the customer using the customerID
	// You can use the customerRepo to fetch the customer from the repository

	// Example:
	customer, err := uc.customerRepo.GetCustomerByID(ctx, customerID)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
