package model

import "time"

type Order struct {
	ID          string `gorm:"primaryKey not null"`
	Items       []*OrderItem
	Total       float64
	Status      string `gorm:"not null"`
	CustomerCPF string
	Customer    Customer
	CreatedAt   time.Time
}

type OrderItem struct {
	ID        string `gorm:"primaryKey not null"`
	ProductID string `gorm:"not null"`
	Product   Product
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	Status    string  `gorm:"not null"`
	OrderID   string
}
