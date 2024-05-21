package portsusecase

import (
	"context"
)

type DeleteProductUseCase interface {
	DeleteProduct(ctx context.Context, id string) error
}
