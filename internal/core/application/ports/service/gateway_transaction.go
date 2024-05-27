package portsusecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type GatewayTransactionService interface {
	CreateCheckout(ctx context.Context, checkout *entity.Checkout) error
}
