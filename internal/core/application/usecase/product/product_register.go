package usecase

import (
	"context"
	"errors"
	"fmt"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type ProductRegisterUseCase struct {
	repository portsrepository.ProductRepository
}

func NewProductRegister(repository portsrepository.ProductRepository) *ProductRegisterUseCase {
	return &ProductRegisterUseCase{
		repository: repository,
	}
}

// RegisterProduct registra um novo cliente.
func (cr *ProductRegisterUseCase) RegisterProduct(ctx context.Context, product entity.Product) error {

	// TODO: aqui tem um problema; o id deve ser gerado no repositorio, então, não faz sentido validar por ID aqui.
	fmt.Println("usecase: verifica se o produto existe: " + product.GetID())
	entityFound, err := cr.repository.Find(ctx, product.GetID())
	if err != nil && err.Error() != "product not found" {
		fmt.Println("usecase: err: " + err.Error())
		return err
	}

	if entityFound != nil {
		fmt.Println("usecase: producto já existe: " + product.GetID())
		return errors.New("product already exists")
	}

	fmt.Println("usecase: Criando produto: " + product.GetID())
	// Cria o cliente
	err = cr.repository.Create(ctx, &product)
	if err != nil {
		return err
	}

	return nil
}
