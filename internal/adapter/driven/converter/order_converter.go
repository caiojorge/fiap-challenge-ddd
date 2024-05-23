package converter

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type OrderConverter struct{}

func NewOrderConverter() *OrderConverter {
	return &OrderConverter{}
}

func (pc *OrderConverter) FromEntity(entity *entity.Order) *model.Order {

	return &model.Order{}
}

// TODO: voltar aqui para avaliar se é melhor retornar um erro tbm
func (pc *OrderConverter) ToEntity(model *model.Order) *entity.Order {
	return &entity.Order{}
}
