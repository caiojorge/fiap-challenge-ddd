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

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()
	return &Server{router: r}
}

func (s *Server) Initialization() {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&model.Customer{})
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	repo := repositorygorm.NewCustomerRepositoryGorm(db)

	uc := usecase.NewCustomerRegister(repo)
	registerController := controller.NewRegisterCustomerController(uc)

	g := s.router.Group("kitchencontrol/api/v1")
	g.GET("/customers", controller.GetCustomerByCPF)
	g.POST("/customers", controller.PostRegisterCustomer)
	g.PUT("/customers/:cpf", controller.PutRegisterCustomer)

}

func (s *Server) GetRouter() *gin.Engine {
	return s.router
}
