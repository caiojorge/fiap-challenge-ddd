package entity

import (
	"errors"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/validator"
)

type Order struct {
	ID          string
	Items       []*OrderItem
	Total       float64
	Status      string
	CustomerCPF string
}

func OrderInit(customerCPF string) *Order {
	order := Order{
		ID:          shared.NewIDGenerator(),
		CustomerCPF: customerCPF,
		Status:      valueobject.OrderStatusConfirmed,
	}

	return &order
}

func NewOrder(cpf string, items []*OrderItem) (*Order, error) {
	order := Order{
		ID:          shared.NewIDGenerator(),
		CustomerCPF: cpf,
		Items:       items,
		Status:      valueobject.OrderStatusConfirmed,
	}

	err := order.Validate()
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *Order) Validate() error {

	if o.CustomerCPF != "" {
		cpfValidator := validator.CPFValidator{}

		err := cpfValidator.Validate(o.CustomerCPF)
		if err != nil {
			return err
		}
	}

	if len(o.Items) == 0 {
		return errors.New("invalid order items")
	}

	return nil
}

func (o *Order) AddItem(item *OrderItem) {
	o.Items = append(o.Items, item)
}

func (o *Order) RemoveItem(item *OrderItem) {
	for i, v := range o.Items {
		if v == item {
			o.Items = append(o.Items[:i], o.Items[i+1:]...)
		}
	}
}

// CalculateTotal se os itens forem confirmados
func (o *Order) CalculateTotal() {
	for _, item := range o.Items {
		if item.Status == valueobject.OrderItemStatusConfirmed {
			o.Total += item.Price
		}
	}
}

func (o *Order) Pay() {
	o.Status = valueobject.OrderStatusPaid
}

func (o *Order) Prepare() {
	o.Status = valueobject.OrderStatusPreparing
}

func (o *Order) Deliver() {
	o.Status = valueobject.OrderStatusDelivered
}

func (o *Order) Cancel() {
	o.Status = valueobject.OrderStatusCanceled
}
