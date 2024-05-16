package converter

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/formatter"
)

func FromEntity(entity *entity.Customer) *model.Customer {

	cpfWithoutNonDigits := formatter.RemoveFormatFromCPF(entity.GetCPF().Value)

	return &model.Customer{
		CPF:   cpfWithoutNonDigits,
		Name:  entity.GetName(),
		Email: entity.GetEmail(),
	}
}

func ToEntity(model *model.Customer) *entity.Customer {
	cpfWithNonDigits, err := formatter.FormatCPF(model.CPF)

	if err != nil {
		return nil
	}

	cpf, _ := valueobject.NewCPF(cpfWithNonDigits)
	customer, _ := entity.NewCustomer(*cpf, model.Name, model.Email)
	return customer
}
