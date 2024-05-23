package server

import (
	"context"
	"log"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/converter"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/repositorygorm"
	controllercustomer "github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/controller/customer"
	controllerorder "github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/controller/order"
	controllerproduct "github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/controller/product"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/infra"
	usecasecustomer "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/usecase/customer"
	usecaseorder "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/usecase/order"
	usecaseproduct "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/usecase/product"
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

	g := s.router.Group("/kitchencontrol/api/v1/customers")
	{
		repo := repositorygorm.NewCustomerRepositoryGorm(db)

		registerController := controllercustomer.NewRegisterCustomerController(ctx, usecasecustomer.NewCustomerRegister(repo))
		g.POST("/", registerController.PostRegisterCustomer)

		updateController := controllercustomer.NewUpdateCustomerController(ctx, usecasecustomer.NewCustomerUpdate(repo))
		g.PUT("/:cpf", updateController.PutUpdateCustomer)

		findByCPFController := controllercustomer.NewFindCustomerByCPFController(ctx, usecasecustomer.NewCustomerFindByCPF(repo))
		g.GET("/:cpf", findByCPFController.GetCustomerByCPF)

		findAllController := controllercustomer.NewFindAllCustomersController(ctx, usecasecustomer.NewCustomerFindAll(repo))
		g.GET("/", findAllController.GetAllCustomers)
	}

	p := s.router.Group("/kitchencontrol/api/v1/products")
	{
		converter := converter.NewProductConverter()
		repo := repositorygorm.NewProductRepositoryGorm(db, converter)

		registerController := controllerproduct.NewRegisterProductController(ctx, usecaseproduct.NewProductRegister(repo))
		p.POST("/", registerController.PostRegisterProduct)

		findAllController := controllerproduct.NewFindAllProductController(ctx, usecaseproduct.NewProductFindAll(repo))
		p.GET("/", findAllController.GetAllProducts)

		findByIDController := controllerproduct.NewFindProductByIDController(ctx, usecaseproduct.NewProductFindByID(repo))
		p.GET("/:id", findByIDController.GetProductByID)

		findByCategoryController := controllerproduct.NewFindProductByCategoryController(ctx, usecaseproduct.NewProductFindByCategory(repo))
		p.GET("/category/:id", findByCategoryController.GetProductByCategory)

		updateController := controllerproduct.NewUpdateProductController(ctx, usecaseproduct.NewProductUpdate(repo))
		p.PUT("/:id", updateController.PutUpdateProduct)

	}

	o := s.router.Group("/kitchencontrol/api/v1/orders")
	{
		converter := converter.NewOrderConverter()
		repo := repositorygorm.NewOrderRepositoryGorm(db, converter)
		orderController := controllerorder.NewCreateOrderController(ctx, usecaseorder.NewOrderCreate(repo))
		o.POST("/", orderController.PostCreateOrder)
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

	// get a connection
	connection := db.GetConnection("mysql")
	if connection == nil {
		log.Fatal("Expected a non-nil MySQL connection, but got nil")
	}

	// Migrate the schema
	if err := connection.AutoMigrate(&model.Customer{}, &model.Product{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
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
