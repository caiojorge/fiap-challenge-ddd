package controller

import (
	"context"
	"net/http"

	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/usecase"
	"github.com/gin-gonic/gin"
)

type FindCustomerByCPFController struct {
	usecase portsusecase.FindCustomerByCPFUseCase
	ctx     context.Context
}

func NewFindCustomerByCPFController(ctx context.Context, usecase portsusecase.FindCustomerByCPFUseCase) *FindCustomerByCPFController {
	return &FindCustomerByCPFController{
		usecase: usecase,
		ctx:     ctx,
	}
}

func (cr *FindCustomerByCPFController) GetCustomerByCPF(c *gin.Context) {
	//cpf, ok := c.GetQuery("cpf")
	cpf := c.Param("cpf")

	if cpf == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	customer, err := cr.usecase.FindCustomerByCPF(cr.ctx, cpf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}
