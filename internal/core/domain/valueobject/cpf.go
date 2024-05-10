package valueobject

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/validator"
)

type CPF struct {
	Value string
}

func NewCPF(value string) (*CPF, error) {

	cpf := &CPF{
		Value: value,
	}

	err := cpf.Validate()
	if err != nil {
		return nil, err
	}

	return cpf, nil
}

func (c *CPF) GetValue() string {
	return c.Value
}

func (c *CPF) Validate() error {
	cpf := c.Value

	cpfValidator := validator.CPFValidator{}

	return cpfValidator.Validate(cpf)
}
