package converter

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
)

func FromEntity(entity *entity.Customer) *model.Customer {
	return &model.Customer{
		CPF:   entity.GetCPF().Value,
		Name:  entity.GetName(),
		Email: entity.GetEmail(),
	}
}

func ToEntity(model *model.Customer) *entity.Customer {
	cpf, _ := valueobject.NewCPF(model.CPF)
	customer, _ := entity.NewCustomer(*cpf, model.Name, model.Email)
	return customer
}
