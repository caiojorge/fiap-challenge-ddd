package entity

import (
	"errors"
	"regexp"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
)

type Customer struct {
	CPF   valueobject.CPF
	Name  string
	Email string
}

func NewCustomer(cpf valueobject.CPF, name, email string) (*Customer, error) {

	customer := &Customer{
		CPF:   cpf,
		Name:  name,
		Email: email,
	}

	// Validate customer
	err := customer.Validate()
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *Customer) GetCPF() valueobject.CPF {
	return c.CPF
}

func (c *Customer) GetName() string {
	return c.Name
}

func (c *Customer) GetEmail() string {
	return c.Email
}

func (c *Customer) Validate() error {

	if c.CPF.GetValue() == "" {
		return errors.New("CPF is required")
	}

	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.Email == "" || !isValidateEmail(c.Email) {
		return errors.New("e-mail is required")
	}

	return nil
}

func isValidateEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}
