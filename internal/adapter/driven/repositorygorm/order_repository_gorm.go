package repositorygorm

import (
	"context"
	"errors"
	"fmt"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/converter"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"gorm.io/gorm"
)

type OrderRepositoryGorm struct {
	DB        *gorm.DB
	converter converter.Converter[entity.Order, model.Order]
}

func NewOrderRepositoryGorm(db *gorm.DB, converter converter.Converter[entity.Order, model.Order]) *OrderRepositoryGorm {
	return &OrderRepositoryGorm{
		DB:        db,
		converter: converter,
	}
}

// Create creates a new product. It returns an error if something goes wrong.
func (r *OrderRepositoryGorm) Create(ctx context.Context, entity *entity.Order) error {
	fmt.Println("repositorygorm: Criando order: " + entity.GetID())
	model := r.converter.FromEntity(entity)

	return r.DB.Create(model).Error
}

func (r *OrderRepositoryGorm) Update(ctx context.Context, entity *entity.Order) error {

	result := r.DB.Save(r.converter.FromEntity(entity))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderRepositoryGorm) Find(ctx context.Context, id string) (*entity.Order, error) {
	var orderModel model.Order
	fmt.Println("repositorygorm: Find order: " + id)
	result := r.DB.Model(&model.Order{}).Where("id = ?", id).First(&orderModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	entity := r.converter.ToEntity(&orderModel)

	return entity, nil
}

func (r *OrderRepositoryGorm) FindAll(ctx context.Context) ([]*entity.Order, error) {
	var mOrders []model.Order
	result := r.DB.Find(&mOrders)
	if result.Error != nil {
		return nil, result.Error
	}

	var eOrders []*entity.Order

	for _, mOrder := range mOrders {
		eOrder := r.converter.ToEntity(&mOrder)
		eOrders = append(eOrders, eOrder)
	}

	return eOrders, nil
}

func (r *OrderRepositoryGorm) Delete(ctx context.Context, id string) error {
	var orderModel model.Order
	result := r.DB.Model(&model.Order{}).Where("id = ?", id).First(&orderModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("repositorygorm: order not found")
			return nil
		}
		return result.Error
	}

	result = r.DB.Delete(&orderModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
