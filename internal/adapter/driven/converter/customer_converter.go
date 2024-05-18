package converter

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/formatter"
)

// TODO: voltar aqui para usar metodos ao invés de func: type CustomerConverter struct{}

func FromEntity(entity *entity.Customer) *model.Customer {

	cpfWithoutNonDigits := formatter.RemoveFormatFromCPF(entity.GetCPF().Value)

	return &model.Customer{
		CPF:   cpfWithoutNonDigits,
		Name:  entity.GetName(),
		Email: entity.GetEmail(),
	}
}

// TODO: voltar aqui para avaliar se é melhor retornar um erro tbm
func ToEntity(model *model.Customer) *entity.Customer {
	cpfWithNonDigits, err := formatter.FormatCPF(model.CPF)

	if err != nil {
		return nil
	}

	cpf, _ := valueobject.NewCPF(cpfWithNonDigits)
	customer, _ := entity.NewCustomer(*cpf, model.Name, model.Email)
	return customer
}
