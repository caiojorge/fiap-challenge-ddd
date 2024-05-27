package usecase

import (
	"context"
	"errors"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

// CheckoutCreateUseCase é a implementação da interface CheckoutCreateUseCase.
// Nesse momento, não iremos implementar a integração com o gateway de pagamento.
type CheckoutCreateUseCase struct {
	orderRepository    portsrepository.OrderRepository
	customerRepository portsrepository.CheckoutRepository
}

func NewCheckoutCreate(orderRepository portsrepository.OrderRepository,
	customerRepository portsrepository.CheckoutRepository) *CheckoutCreateUseCase {
	return &CheckoutCreateUseCase{
		orderRepository:    orderRepository,
		customerRepository: customerRepository,
	}
}

// CreateCheckout registra um novo pedido.
func (cr *CheckoutCreateUseCase) CreateCheckout(ctx context.Context, checkout *entity.Checkout) error {

	order, err := cr.orderRepository.Find(ctx, checkout.OrderID)
	if err != nil {
		return err
	}

	if order == nil {
		return errors.New("order not found")
	}

	cupon := 0.0
	// o valor total do pagamento é o total do pedido.
	// poderiamos aplicar, por exemplo, um desconto ou acrescimo. (pedido - cupom ou pedido + frete)
	checkout.FinalAmount(order.Total - cupon)

	// validação dos dados do pagamento
	err = checkout.Validate()
	if err != nil {
		return err
	}

	err = cr.customerRepository.Create(ctx, checkout)
	if err != nil {
		return err
	}

	return nil
}
