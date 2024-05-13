package server

import (
	"context"
	"log"
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

func (s *GinServer) Initialization() *GinServer {

	db := setupDB()
	ctx := setupContext()

	g := s.router.Group("kitchencontrol/api/v1")

	setupCustomerRoutes(ctx, db, g)

	return s
}

func setupCustomerRoutes(ctx context.Context, db *gorm.DB, g *gin.RouterGroup) {
	repo := repositorygorm.NewCustomerRepositoryGorm(db)

	uc := usecase.NewCustomerRegister(repo)
	registerController := controller.NewRegisterCustomerController(ctx, uc)
	g.POST("/customers", registerController.PostRegisterCustomer)
	g.PUT("/customers/:cpf", registerController.PutRegisterCustomer)

	findByCPFController := controller.NewFindCustomerByCPFController(ctx, usecase.NewCustomerFindByCPF(repo))
	g.GET("/customers/cpf", findByCPFController.GetCustomerByCPF)

	findAllController := controller.NewFindAllCustomersController(ctx, usecase.NewCustomerFindAll(repo))
	g.GET("/customers", findAllController.GetAllCustomers)

}

func setupContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	return ctx
}

func setupDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Customer{})
	return db
}

func (s *GinServer) Run(port string) {
	if err := s.router.Run(port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func (s *GinServer) GetRouter() *gin.Engine {
	return s.router
}
