package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomerByCPF(c *gin.Context) {
	cpf, ok := c.GetQuery("cpf")
	if !ok {
		// Handle the case where the id is missing.
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing cpf"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile", "cpf": cpf})
}
