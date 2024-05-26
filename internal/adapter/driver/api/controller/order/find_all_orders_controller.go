package controller

import (
	"context"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/usecase/order"
	"github.com/gin-gonic/gin"
)

type FindAllController struct {
	usecase portsusecase.CreateOrderUseCase
	ctx     context.Context
}

func NewFindAllController(ctx context.Context, usecase portsusecase.CreateOrderUseCase) *FindAllController {
	return &FindAllController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// GetAllOrders returns a list of all orders
// @Summary Get all orders
// @Description Get details of all orders
// @Tags Orders
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.OrderDTO
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /customers [get]
func (r *FindAllController) GetAllOrders(c *gin.Context) {

	customers, err := r.usecase.FindAllCustomers(cr.ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if len(customers) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No customers found"})
		return
	}

	var dtoCustomers []dto.CustomerDTO
	dto := dto.CustomerDTO{}
	for _, customer := range customers {
		dto.FromEntity(*customer)
		dtoCustomers = append(dtoCustomers, dto)
	}

	c.JSON(http.StatusOK, dtoCustomers)
}
