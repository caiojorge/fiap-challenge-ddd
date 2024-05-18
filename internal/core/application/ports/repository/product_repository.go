package portsrepository

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

// ProductRepository defines the methods for interacting with the product data.
type ProductRepository interface {
	Create(ctx context.Context, product *entity.Product) error
	Update(ctx context.Context, product *entity.Product) error
	Find(ctx context.Context, id string) (*entity.Product, error)
	FindAll(ctx context.Context) ([]*entity.Product, error)
	Delete(ctx context.Context, id string) error
}
