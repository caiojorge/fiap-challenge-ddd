package usecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type CustomerRegister struct {
	repository repository.CustomerRepository
}

func NewCustomerRegister(repository repository.CustomerRepository) *CustomerRegister {
	return &CustomerRegister{
		repository: repository,
	}
}

// RegisterCustomer registra um novo cliente.
func (cr *CustomerRegister) RegisterCustomer(ctx context.Context, customer entity.Customer) error {

	_, err := cr.repository.Find(ctx, customer.GetCPF().Value)
	if err != nil && err.Error() != "customer not found" {
		return err
	}

	err = cr.repository.Create(ctx, &customer)
	if err != nil {
		return err
	}

	return nil
}
