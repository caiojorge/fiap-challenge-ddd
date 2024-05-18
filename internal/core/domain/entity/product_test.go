package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {

	product, err := NewProduct("Lanche XPTO", "Pão queijo e carne", "Lanche", 30.00)
	assert.Nil(t, err)
	assert.NotNil(t, product)

	assert.Equal(t, "Lanche XPTO", product.Name)

	assert.NotEmpty(t, product.GetID())

}
