package service

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared"
)

// PaymentService provides methods for payment operations.
type PaymentService struct {
}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

// CreateCheckout creates a new checkout. This method should be implemented by the payment gateway.
// Fake implementation.
func (p *PaymentService) CreateCheckout(ctx context.Context, checkout *entity.Checkout) (*string, error) {
	transactionID := shared.NewIDGenerator() // Fake transaction ID
	return &transactionID, nil
}
