package controller

import (
	"context"
	"net/http"

	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/usecase"
	"github.com/gin-gonic/gin"
)

type FindAllCustomersController struct {
	usecase portsusecase.FindAllCustomersUseCase
	ctx     context.Context
}

func NewFindAllCustomersController(ctx context.Context, usecase portsusecase.FindAllCustomersUseCase) *FindAllCustomersController {
	return &FindAllCustomersController{
		usecase: usecase,
		ctx:     ctx,
	}
}

func (cr *FindAllCustomersController) GetAllCustomers(c *gin.Context) {

	customers, err := cr.usecase.FindAllCustomers(cr.ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	c.JSON(http.StatusOK, customers)
}
