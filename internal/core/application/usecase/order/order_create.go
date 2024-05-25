package usecase

import (
	"context"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type OrderCreateUseCase struct {
	repository portsrepository.OrderRepository
}

func NewOrderCreate(repository portsrepository.OrderRepository) *OrderCreateUseCase {
	return &OrderCreateUseCase{
		repository: repository,
	}
}

// CreateOrder registra um novo cliente.
func (cr *OrderCreateUseCase) CreateOrder(ctx context.Context, order *entity.Order) error {

	order.ConfirmOrder()

	err := cr.repository.Create(ctx, order)
	if err != nil {
		return err
	}

	return nil
}
