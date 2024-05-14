package model

import (
	"fmt"

	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/validator"
)

// Customer representa um cliente no banco de dados.
type Customer struct {
	CPF   string `gorm:"primaryKey"`
	Name  string
	Email string
}

// Validate verifica se os campos obrigatórios de um cliente estão preenchidos.
func (c *Customer) Validate() error {

	cpfValidator := validator.CPFValidator{}
	if cpfValidator.Validate(c.CPF) != nil {
		return fmt.Errorf("cpf is invalid")
	}

	if c.Name != "" && len(c.Name) < 3 {
		return fmt.Errorf("name should have at least 3 characters")
	}

	emailValidator := validator.EmailValidator{}
	if c.Email != "" && !emailValidator.IsValid(c.Email) {
		return fmt.Errorf("email is invalid")
	}

	return nil
}
