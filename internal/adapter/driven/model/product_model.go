package model

import (
	"fmt"
)

// Product representa um producto no banco de dados.
type Product struct {
	ID          string  `gorm:"primaryKey not null"`
	Name        string  `gorm:"not null;unique;index:idx_name_category"`
	Description string  `gorm:"not null"`
	Category    string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
}

// Validate verifica se os campos obrigatórios de um producto estão preenchidos.
func (c *Product) Validate() error {

	if c.ID == "" {
		return fmt.Errorf("id is required")
	}

	if c.Name == "" {
		return fmt.Errorf("name is required")
	}

	if c.Description == "" {
		return fmt.Errorf("description is required")
	}

	if c.Category == "" {
		return fmt.Errorf("category is required")
	}

	if c.Price == 0 {
		return fmt.Errorf("price is required")
	}

	return nil
}
