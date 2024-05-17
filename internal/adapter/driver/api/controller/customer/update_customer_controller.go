package controller

import (
	"context"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/usecase"
	"github.com/gin-gonic/gin"
)

type UpdateCustomerController struct {
	usecase portsusecase.UpdateCustomerUseCase
	ctx     context.Context
}

func NewUpdateCustomerController(ctx context.Context, usecase portsusecase.UpdateCustomerUseCase) *UpdateCustomerController {
	return &UpdateCustomerController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PutUpdateCustomer updates a customer by cpf
// @Summary Update a customer
// @Description Update details of a customer by cpf
// @Tags Customers
// @Accept  json
// @Produce  json
// @Param id path string true "Customer cpf"
// @Param customer body dto.CustomerDTO true "Customer data"
// @Success 200 {object} dto.CustomerDTO
// @Failure 400 {object} map[string]string "Invalid data"
// @Failure 404 {object} map[string]string "Customer not found"
// @Router /customers/{cpf} [put]
func (r *UpdateCustomerController) PutUpdateCustomer(c *gin.Context) {
	cpf := c.Param("cpf")
	if cpf == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	var dto dto.CustomerDTO

	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	customer, err := dto.ToEntity()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err = r.usecase.UpdateCustomer(r.ctx, *customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto)
}
