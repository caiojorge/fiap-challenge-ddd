package controller

import (
	"context"
	"fmt"
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

// PostRegisterCustomer godoc
// @Summary Create Customer
// @Schemes
// @Description Create Customer in DB
// @Tags Customers
// @Accept json
// @Produce json
// @Param        request   body     dto.CustomerDTO  true  "user request"
// @Success 200 {object} dto.CustomerDTO
// @Router /customers [post]
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

	fmt.Println("controller: Criando cliente: " + dto.CPF)
	r.usecase.RegisterCustomer(r.ctx, *entity)

	// Use the user object, e.g., save to database, etc.
	// gin.H{"status": "customer created " + dto.Name}
	c.JSON(http.StatusOK, dto)
}
