package usecase

import (
	"context"
	"errors"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type ProductUpdateUseCase struct {
	repository portsrepository.ProductRepository
}

func NewProductUpdate(repository portsrepository.ProductRepository) *ProductUpdateUseCase {
	return &ProductUpdateUseCase{
		repository: repository,
	}
}

// UpdateProduct atualiza um novo produto.
func (cr *ProductUpdateUseCase) UpdateProduct(ctx context.Context, product entity.Product) error {

	prd, err := cr.repository.Find(ctx, product.GetID())
	if err != nil {
		return err
	}

	if prd == nil {
		return errors.New("product not found")
	}

	err = cr.repository.Update(ctx, &product)
	if err != nil {
		return err
	}

	return nil
}
