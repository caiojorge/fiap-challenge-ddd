package controller

import (
	"context"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/usecase/order"
	"github.com/gin-gonic/gin"
)

type FindAllController struct {
	usecase portsusecase.FindAllOrderUseCase
	ctx     context.Context
}

func NewFindAllController(ctx context.Context, usecase portsusecase.FindAllOrderUseCase) *FindAllController {
	return &FindAllController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// GetAllOrders returns a list of all orders
// @Summary Get all orders
// @Description Retorna todos os pedidos (orders) registrados no sistema. Se não houver pedidos, retorna um erro (404).
// @Tags Orders
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.OrderDTO
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Router /orders [get]
func (r *FindAllController) GetAllOrders(c *gin.Context) {

	orders, err := r.usecase.FindAllOrders(r.ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No orders found"})
		return
	}

	var dtos []dto.OrderDTO
	dto := dto.OrderDTO{}
	for _, order := range orders {
		dto.FromEntity(*order)
		dtos = append(dtos, dto)
	}

	// err = copier.Copy(&dtos, &orders)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
	// 	return
	// }

	c.JSON(http.StatusOK, dtos)
}
