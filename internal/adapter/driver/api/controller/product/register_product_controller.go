package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"
	"github.com/gin-gonic/gin"
)

var ErrCustomerAlreadyExists = errors.New("customer already exists")

type RegisterProductController struct {
	//usecase portsusecase.RegisterCustomerUseCase
	ctx context.Context
}

func NewRegisterProductController(ctx context.Context) *RegisterProductController {
	return &RegisterProductController{
		//usecase: usecase,
		ctx: ctx,
	}
}

// PostRegisterProduct godoc
// @Summary Create Product
// @Schemes
// @Description Create Product in DB
// @Tags Products
// @Accept json
// @Produce json
// @Param        request   body     dto.ProductDTO  true  "user request"
// @Success 200 {object} dto.ProductDTO
// @Failure 400 {object} string
// @Router /customers [post]
func (r *RegisterProductController) PostRegisterProduct(c *gin.Context) {
	var dto dto.ProductDTO

	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// entity, err := dto.ToEntity()
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
	// 	return
	// }

	// fmt.Println("controller: Criando cliente: " + dto.CPF)
	// err = r.usecase.RegisterCustomer(r.ctx, *entity)
	// if err != nil {
	// 	if err == ErrCustomerAlreadyExists {
	// 		c.JSON(http.StatusConflict, gin.H{"error": "Customer already exists"})
	// 	} else {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	}
	// 	return
	// }

	c.JSON(http.StatusOK, dto)
}
