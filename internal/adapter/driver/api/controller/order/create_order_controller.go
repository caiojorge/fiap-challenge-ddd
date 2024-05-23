package controller

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/usecase/order"
	"github.com/gin-gonic/gin"
)

var ErrAlreadyExists = errors.New("order already exists")

type CreateOrderController struct {
	usecase portsusecase.CreateOrderUseCase
	ctx     context.Context
}

func NewCreateOrderController(ctx context.Context, usecase portsusecase.CreateOrderUseCase) *CreateOrderController {
	return &CreateOrderController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PostCreateOrder godoc
// @Summary Create Order
// @Schemes
// @Description Create Order in DB
// @Tags Orders
// @Accept json
// @Produce json
// @Param        request   body     dto.OrderDTO  true  "cria nova Order"
// @Success 200 {object} dto.OrderDTO
// @Failure 400 {object} string "invalid data"
// @Failure 409 {object} string "Order already exists"
// @Failure 500 {object} string "internal server error"
// @Router /orders [post]
func (r *CreateOrderController) PostCreateOrder(c *gin.Context) {
	var dto dto.OrderDTO

	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	entity, err := dto.ToEntity()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Nesse cenário, o ID informado será ignorado e um novo ID será gerado
	fmt.Println("controller: Criando Order: " + entity.CustomerCPF + " - " + entity.ID)

	err = r.usecase.CreateOrder(r.ctx, entity)
	if err != nil {
		if err == ErrAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": "Order already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	dto.FromEntity(*entity)

	log.Println("Order created: ", dto.ID)

	c.JSON(http.StatusOK, dto)
}
