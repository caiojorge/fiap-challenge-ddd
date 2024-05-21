package controller

import (
	"context"
	"net/http"

	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/usecase/product"
	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	usecase portsusecase.DeleteProductUseCase
	ctx     context.Context
}

func NewDeleteProductController(ctx context.Context, usecase portsusecase.DeleteProductUseCase) *DeleteProductController {
	return &DeleteProductController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// DeleteProduct updates a Product by id
// @Summary Update a Product
// @Description Update details of a Product by id
// @Tags Products
// @Accept  json
// @Produce  json
// @Param id path string true "Product id"
// @Param Product body dto.ProductDTO true "Product data"
// @Success 200 {object} dto.ProductDTO
// @Failure 400 {object} string "Invalid data"
// @Failure 404 {object} string "Product not found"
// @Router /products/{id} [delete]
func (r *DeleteProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err := r.usecase.DeleteProduct(r.ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}
