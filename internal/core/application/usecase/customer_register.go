package usecase

import (
	"context"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type CustomerRegisterUseCase struct {
	repository ports.CustomerRepository
}

func NewCustomerRegister(repository ports.CustomerRepository) *CustomerRegisterUseCase {
	return &CustomerRegisterUseCase{
		repository: repository,
	}
}

// RegisterCustomer registra um novo cliente.
func (cr *CustomerRegisterUseCase) RegisterCustomer(ctx context.Context, customer entity.Customer) error {

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
