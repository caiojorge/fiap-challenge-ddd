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

// DeleteProduct deletes a Product by id
// @Summary Delete a Product
// @Description Delete details of a Product by id
// @Tags Products
// @Accept  json
// @Produce  json
// @Param id path string true "Product id"
// @Success 200 {object} dto.ProductDTO
// @Failure 400 {object} string "Invalid data"
// @Router /products/{id} [delete]
func (r *DeleteProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err := r.usecase.DeleteProduct(r.ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": id + " deleted"})
}
