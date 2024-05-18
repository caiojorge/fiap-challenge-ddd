package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Category    string
}

func NewProduct(name, description, category string, price float64) (*Product, error) {

	product := &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {

	if p.Name == "" {
		return errors.New("name is required")
	}

	if p.Description == "" {
		return errors.New("description is required")
	}

	if p.Price == 0 {
		return errors.New("price is required")
	}

	if p.Category == "" {
		return errors.New("category is required")
	}

	return nil
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetDescription() string {
	return p.Description
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetCategory() string {
	return p.Category
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) RedifneID(id string) {
	p.ID = id
}

func (p *Product) ChangePrice(price float64) {
	p.Price = price
}