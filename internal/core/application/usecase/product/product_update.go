package usecase

import (
	"context"
	"fmt"

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

	fmt.Println("usecase: verifica se o product existe: " + product.GetID())
	_, err := cr.repository.Find(ctx, product.GetID())
	if err != nil && err.Error() != "product not found" {
		fmt.Println("usecase: err: " + err.Error())
		return err
	}

	fmt.Println("usecase: atualizando product: " + product.GetID())
	err = cr.repository.Update(ctx, &product)
	if err != nil {
		return err
	}

	return nil
}
