package usecase

import (
	"context"
	"errors"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/formatter"
)

// OrderCreateUseCase é a implementação de OrderCreateUseCase.
// Um caso de uso é uma estrutura que contém todas as regras de negócio para uma determinada funcionalidade.
// Nesse cenário, vamos precisar acessar 3 agregados e seus repositorios.
type OrderCreateUseCase struct {
	orderRepository    portsrepository.OrderRepository
	customerRepository portsrepository.CustomerRepository
	productRepository  portsrepository.ProductRepository
}

func NewOrderCreate(orderRepository portsrepository.OrderRepository,
	customerRepository portsrepository.CustomerRepository,
	productRepository portsrepository.ProductRepository) *OrderCreateUseCase {
	return &OrderCreateUseCase{
		orderRepository:    orderRepository,
		customerRepository: customerRepository,
		productRepository:  productRepository,
	}
}

// CreateOrder registra um novo pedido.
func (cr *OrderCreateUseCase) CreateOrder(ctx context.Context, order *entity.Order) error {

	// se o cpf for empty, indica que o cliente não quis se identificar, e isso esta ok, segundo as regras de negócio
	if order.CustomerCPF != "" {
		// busca o cliente pelo cpf
		customer, err := cr.customerRepository.Find(ctx, order.CustomerCPF)
		if err != nil {
			return err
		}

		// se o cliente for informado, temos q validar o cpf, e cadastra-lo caso não exista
		// só entra aqui se o cliente não existir na base de dados
		// se o cliente não for nulo, ele já existe o cpf é considerado válido
		// em teoria, não existe ordem duplicada. o mesmo cliente pode comprar várias vezes. (não vou validar isso aqui)
		if customer == nil {
			// o cliente não é obrigatório, mas se for informado, ele precisa ser válido.
			// apenas nesse caso, se o cliente não existir, ele será persistido. (apenas o cpf)
			cpf, err := valueobject.NewCPF(order.CustomerCPF)
			if err != nil {
				return err
			}

			// identifica o cliente pelo cpf
			newCustomer, err := entity.NewCustomerWithCPFOnly(cpf)
			if err != nil {
				return err
			}

			// cria o cliente sem o nome e email
			err = cr.customerRepository.Create(ctx, newCustomer)
			if err != nil {
				return err
			}
		}

		// just in case... remove a máscara do cpf
		order.CustomerCPF = formatter.RemoveMaksFromCPF(order.CustomerCPF)
	}

	// valida se os produtos informados existem
	for _, item := range order.Items {
		product, err := cr.productRepository.Find(ctx, item.ProductID)
		if err != nil {
			return err
		}

		if product == nil {
			return errors.New("product not found")
		}

		// atualiza o preço do produto
		item.UpdatePrice(product.Price)
	}

	// toda regra de negócio para criar uma ordem confirmada
	order.ConfirmOrder()

	// cria a ordem e usa o cliente (novo ou existente) e o produto existente.
	err := cr.orderRepository.Create(ctx, order)
	if err != nil {
		return err
	}

	return nil
}
