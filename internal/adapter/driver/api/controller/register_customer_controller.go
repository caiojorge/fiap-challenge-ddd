package controller

import (
	"context"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/usecase"
	"github.com/gin-gonic/gin"
)

type RegisterCustomerController struct {
	usecase portsusecase.RegisterCustomerUseCase
	ctx     context.Context
}

func NewRegisterCustomerController(ctx context.Context, usecase portsusecase.RegisterCustomerUseCase) *RegisterCustomerController {
	return &RegisterCustomerController{
		usecase: usecase,
		ctx:     ctx,
	}
}

func (r *RegisterCustomerController) PostRegisterCustomer(c *gin.Context) {
	var dto dto.CustomerDTO

	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	entity, err := dto.ToEntity()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	r.usecase.RegisterCustomer(r.ctx, *entity)

	// Use the user object, e.g., save to database, etc.
	c.JSON(http.StatusOK, gin.H{"status": "customer created " + dto.Name})
}

func (r *RegisterCustomerController) PutRegisterCustomer(c *gin.Context) {
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
	// Use the user object, e.g., save to database, etc.
	c.JSON(http.StatusOK, gin.H{"status": "customer created " + dto.Name})
}