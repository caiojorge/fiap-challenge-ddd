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

type CustomerRepositoryGorm struct {
	DB *gorm.DB
}

func NewCustomerRepositoryGorm(db *gorm.DB) *CustomerRepositoryGorm {
	return &CustomerRepositoryGorm{
		DB: db,
	}
}

func (r *CustomerRepositoryGorm) Create(ctx context.Context, entity *entity.Customer) error {
	fmt.Println("repositorygorm: Criando cliente: " + entity.GetCPF().Value)
	model := converter.FromEntity(entity)
	return r.DB.Create(model).Error
}

func (r *CustomerRepositoryGorm) Update(ctx context.Context, entity *entity.Customer) error {

	result := r.DB.Model(&model.Customer{}).Where("cpf = ?", entity.GetCPF().Value).Updates(model.Customer{Name: entity.GetName(), Email: entity.GetEmail()})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CustomerRepositoryGorm) Find(ctx context.Context, id string) (*entity.Customer, error) {
	var customerModel model.Customer
	fmt.Println("repositorygorm: Find cliente: " + id)
	result := r.DB.Model(&model.Customer{}).Where("cpf = ?", id).First(&customerModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	entity := converter.ToEntity(&customerModel)

	return entity, nil
}

func (r *CustomerRepositoryGorm) FindAll(ctx context.Context) ([]*entity.Customer, error) {
	var mCustomers []model.Customer
	result := r.DB.Find(&mCustomers)
	if result.Error != nil {
		return nil, result.Error
	}

	var eCustomers []*entity.Customer

	for _, mCustomer := range mCustomers {
		eCustomer := converter.ToEntity(&mCustomer)
		eCustomers = append(eCustomers, eCustomer)
	}

	return eCustomers, nil
}
