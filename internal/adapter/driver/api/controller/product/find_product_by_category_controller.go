package controller

import (
	"context"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/usecase/product"
	"github.com/gin-gonic/gin"
)

type FindProductByCategoryController struct {
	usecase portsusecase.FindProductByCategoryUseCase
	ctx     context.Context
}

func NewFindProductByCategoryController(ctx context.Context, usecase portsusecase.FindProductByCategoryUseCase) *FindProductByCategoryController {
	return &FindProductByCategoryController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// @Summary Get a Product
// @Description Get details of a Product by id
// @Tags Products
// @Accept  json
// @Produce  json
// @Param id path string true "Product id"
// @Success 200 {object} dto.ProductDTO
// @Failure 404 {object} string "Product not found"
// @Failure 500 {object} string "Product not found"
// @Router /products/{id} [get]
func (cr *FindProductByCategoryController) GetProductByCategory(c *gin.Context) {
	//id, ok := c.GetQuery("id")
	category := c.Param("category")

	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	product, err := cr.usecase.FindProductByCategory(cr.ctx, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dto := dto.ProductDTO{}
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	dto.FromEntity(*product)
	c.JSON(http.StatusOK, dto)
}
