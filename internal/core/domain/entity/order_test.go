package entity

import (
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/validator"
	"github.com/stretchr/testify/assert"
)

// Test Order

func TestOrder(t *testing.T) {
	// Customer
	cpf, err := valueobject.NewCPF("19528476562")
	assert.Nil(t, err)
	customer, err := NewCustomer(*cpf, "Caio", "email@email.com")
	assert.Nil(t, err)
	assert.Equal(t, "Caio", customer.Name)

	// Product
	product, err := NewProduct("Lanche xpto", "Pão, carne e queijo", "lanche", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)

	// Order
	order := OrderInit(customer.GetCPF().Value)
	order.ID = "1"
	assert.Equal(t, "1", order.ID)
	assert.Equal(t, "19528476562", order.CustomerCPF)

	//orderItem := NewOrderItem(product.GetID(), 1, 10)

	// validando o cpf q seria informado pelo cliente.
	v := validator.CPFValidator{}
	assert.True(t, v.IsValid(order.CustomerCPF))
	err = v.Validate(order.CustomerCPF)
	assert.Nil(t, err)

	err = order.Validate() // não tem itens, e vai dar erro.
	assert.NotNil(t, err)
}
