package dto

import "time"

type CreateCheckoutDTO struct {
	OrderID     string    `json:"order_id"`
	Gateway     string    `json:"gateway"`
	GatewayID   string    `json:"gateway_id"`
	CustomerCPF string    `json:"customer_cpf"`
	CreatedAt   time.Time `json:"created_at"`
}

type CheckoutDTO struct {
	ID          string    `json:"id"`
	OrderID     string    `json:"order_id"`
	Gateway     string    `json:"gateway"`
	GatewayID   string    `json:"gateway_id"`
	CustomerCPF string    `json:"customer_cpf"`
	CreatedAt   time.Time `json:"created_at"`
}
