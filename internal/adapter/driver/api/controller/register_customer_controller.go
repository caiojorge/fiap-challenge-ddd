package controller

import (
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/usecase"
	"github.com/gin-gonic/gin"
)

type RegisterCustomerController struct {
	usecase portsusecase.RegisterCustomerUseCase
}

func NewRegisterCustomerController(usecase portsusecase.RegisterCustomerUseCase) *RegisterCustomerController {
	return &RegisterCustomerController{
		usecase: usecase,
	}
}

func PostRegisterCustomer(c *gin.Context) {
	var dto dto.CustomerDTO

	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	// Use the user object, e.g., save to database, etc.
	c.JSON(http.StatusOK, gin.H{"status": "customer created " + dto.Name})
}

func PutRegisterCustomer(c *gin.Context) {
	cpf := c.Param("cpf")
	if cpf == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	var dto dto.CustomerDTO

	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	// Use the user object, e.g., save to database, etc.
	c.JSON(http.StatusOK, gin.H{"status": "customer created " + dto.Name})
}
