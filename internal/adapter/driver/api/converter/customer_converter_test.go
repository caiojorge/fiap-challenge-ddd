package converter_test

import (
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/converter"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestFromEntity(t *testing.T) {
	// TODO: Add test cases for FromEntity function

	cpf, err := valueobject.NewCPF("123.456.789-09")
	assert.NotNil(t, cpf)
	assert.Nil(t, err)
	customer, err := entity.NewCustomer(*cpf, "John Doe", "email@email.com")
	assert.NotNil(t, customer)
	assert.Nil(t, err)

	//var model *model.Customer
	model := converter.FromEntity(customer)
	assert.NotNil(t, model)
	assert.Equal(t, customer.GetCPF().Value, model.CPF)
	assert.Equal(t, customer.GetName(), model.Name)
	assert.Equal(t, customer.GetEmail(), model.Email)

}

func TestToEntity(t *testing.T) {
	// TODO: Add test cases for ToEntity function
}
