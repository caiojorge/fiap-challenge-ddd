package usecase

import (
	"context"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type ProductFindByIDUseCase struct {
	repository ports.ProductRepository
}

func NewProductFindByID(repository ports.ProductRepository) *ProductFindByIDUseCase {
	return &ProductFindByIDUseCase{
		repository: repository,
	}
}

func (cr *ProductFindByIDUseCase) FindProductByID(ctx context.Context, id string) (*entity.Product, error) {

	product, err := cr.repository.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
