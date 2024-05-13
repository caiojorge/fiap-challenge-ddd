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
	cpf, ok := c.GetQuery("cpf")
	if !ok {
		// Handle the case where the id is missing.
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing cpf"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile", "cpf": cpf})
}
