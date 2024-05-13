package server

import (
	"context"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/repositorygorm"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/controller"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/application/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GinServer struct {
	router *gin.Engine
}

func NewServer() *GinServer {
	r := gin.Default()
	return &GinServer{router: r}
}

func (s *GinServer) Initialization() {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Customer{})
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	repo := repositorygorm.NewCustomerRepositoryGorm(db)

	uc := usecase.NewCustomerRegister(repo)
	registerController := controller.NewRegisterCustomerController(ctx, uc)

	g := s.router.Group("kitchencontrol/api/v1")
	g.GET("/customers", controller.GetCustomerByCPF)
	g.POST("/customers", registerController.PostRegisterCustomer)
	g.PUT("/customers/:cpf", registerController.PutRegisterCustomer)

}

func (s *GinServer) GetRouter() *gin.Engine {
	return s.router
}
