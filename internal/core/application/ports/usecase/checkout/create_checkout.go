package portsusecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type CreateCheckoutUseCase interface {
	CreateCheckout(ctx context.Context, order *entity.Checkout) error
}
