package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
)

func TestNewCustomer(t *testing.T) {

	cpf, err := valueobject.NewCPF("123.456.789-09")
	assert.Nil(t, err)
	assert.NotNil(t, cpf)

	customer, err := NewCustomer(*cpf, "John Doe", "email@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, customer)

}

func TestCustomer_GetCPF(t *testing.T) {
	// TODO: Add test cases for GetCPF method

	cpf, err := valueobject.NewCPF("123.456.789-09")
	assert.Nil(t, err)
	assert.NotNil(t, cpf)

	customer, err := NewCustomer(*cpf, "John Doe", "email@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, customer)

	assert.Equal(t, cpf, customer.GetCPF())

}

func TestCustomer_GetName(t *testing.T) {

	cpf, err := valueobject.NewCPF("123.456.789-09")
	assert.Nil(t, err)
	assert.NotNil(t, cpf)

	customer, err := NewCustomer(*cpf, "John Doe", "email@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, customer)

	assert.Equal(t, "John Doe", customer.GetName())

}

func TestCustomer_GetEmail(t *testing.T) {
	cpf, err := valueobject.NewCPF("123.456.789-09")
	assert.Nil(t, err)
	assert.NotNil(t, cpf)

	customer, err := NewCustomer(*cpf, "John Doe", "email@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, customer)

	assert.Equal(t, "email@email.com", customer.GetEmail())
}

func TestIsValidEmail(t *testing.T) {
	// TODO: Add test cases for isValidateEmail function

	assert.True(t, isValidateEmail("email@email.com"))
	assert.False(t, isValidateEmail("email"))
	assert.False(t, isValidateEmail("email@"))
	assert.False(t, isValidateEmail("email.com"))
	assert.False(t, isValidateEmail("email@.com"))
	assert.False(t, isValidateEmail("email@com"))

}
