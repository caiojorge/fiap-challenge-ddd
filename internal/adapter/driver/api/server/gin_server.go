package server

import (
	"context"
	"log"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/repositorygorm"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/controller"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/infra"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/application/usecase"
	"github.com/gin-gonic/gin"
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

	//db := setupSQLite()
	db := setupDB()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	g := s.router.Group("/kitchencontrol/api/v1")
	{
		repo := repositorygorm.NewCustomerRepositoryGorm(db)

		uc := usecase.NewCustomerRegister(repo)
		registerController := controller.NewRegisterCustomerController(ctx, uc)
		g.POST("/customers", registerController.PostRegisterCustomer)

		updateController := controller.NewUpdateCustomerController(ctx, usecase.NewCustomerUpdate(repo))
		g.PUT("/customers/:cpf", updateController.PutUpdateCustomer)

		findByCPFController := controller.NewFindCustomerByCPFController(ctx, usecase.NewCustomerFindByCPF(repo))
		g.GET("/customers/:cpf", findByCPFController.GetCustomerByCPF)

		findAllController := controller.NewFindAllCustomersController(ctx, usecase.NewCustomerFindAll(repo))
		g.GET("/customers", findAllController.GetAllCustomers)
	}

	return s
}

func setupDB() *gorm.DB {

	host := "localhost"
	port := "3306"
	user := "root"
	password := "root"
	dbName := "dbcontrol"

	db := infra.NewDB(host, port, user, password, dbName)

	connection := db.GetConnection("mysql")
	if connection == nil {
		log.Fatal("Expected a non-nil MySQL connection, but got nil")
	}

	return connection
}

func (s *GinServer) Run(port string) {
	if err := s.router.Run(port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func (s *GinServer) GetRouter() *gin.Engine {
	return s.router
}
