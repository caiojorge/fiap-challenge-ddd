package usecase

import (
	"context"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type CustomerRegisterUseCase struct {
	repository portsrepository.CustomerRepository
}

func NewCustomerRegister(repository portsrepository.CustomerRepository) *CustomerRegisterUseCase {
	return &CustomerRegisterUseCase{
		repository: repository,
	}
}

// RegisterCustomer registra um novo cliente.
func (cr *CustomerRegisterUseCase) RegisterCustomer(ctx context.Context, customer entity.Customer) error {

	// Verifica se o cliente já existe
	_, err := cr.repository.Find(ctx, customer.GetCPF().Value)
	if err != nil && err.Error() != "customer not found" {
		return err
	}

	// Cria o cliente
	err = cr.repository.Create(ctx, &customer)
	if err != nil {
		return err
	}

	return nil
}
