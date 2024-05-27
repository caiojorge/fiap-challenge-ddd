package entity

type Checkout struct {
	ID          string
	OrderID     string
	Gateway     string
	GatewayID   string
	CustomerCPF string
}
