package model

import "time"

type Order struct {
	ID          string
	Items       []*OrderItem
	Total       float64
	Status      string
	CustomerCPF string
	CreatedAt   time.Time
}

type OrderItem struct {
	ID        string
	ProductID string
	Quantity  int
	Price     float64
	Status    string
}
