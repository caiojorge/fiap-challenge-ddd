package server

import (
	"context"
	"log"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/converter"
	infra "github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/db"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/repositorygorm"
	controllercustomer "github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/controller/customer"
	controllerorder "github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/controller/order"
	controllerproduct "github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/controller/product"
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

	customerRepo := repositorygorm.NewCustomerRepositoryGorm(db)
	productConverter := converter.NewProductConverter()
	productRepo := repositorygorm.NewProductRepositoryGorm(db, productConverter)
	orderConverter := converter.NewOrderConverter()
	orderRepo := repositorygorm.NewOrderRepositoryGorm(db, orderConverter)

	g := s.router.Group("/kitchencontrol/api/v1/customers")
	{
		registerController := controllercustomer.NewRegisterCustomerController(ctx, usecasecustomer.NewCustomerRegister(customerRepo))
		g.POST("/", registerController.PostRegisterCustomer)

		updateController := controllercustomer.NewUpdateCustomerController(ctx, usecasecustomer.NewCustomerUpdate(customerRepo))
		g.PUT("/:cpf", updateController.PutUpdateCustomer)

		findByCPFController := controllercustomer.NewFindCustomerByCPFController(ctx, usecasecustomer.NewCustomerFindByCPF(customerRepo))
		g.GET("/:cpf", findByCPFController.GetCustomerByCPF)

		findAllController := controllercustomer.NewFindAllCustomersController(ctx, usecasecustomer.NewCustomerFindAll(customerRepo))
		g.GET("/", findAllController.GetAllCustomers)
	}

	p := s.router.Group("/kitchencontrol/api/v1/products")
	{
		registerController := controllerproduct.NewRegisterProductController(ctx, usecaseproduct.NewProductRegister(productRepo))
		p.POST("/", registerController.PostRegisterProduct)

		findAllController := controllerproduct.NewFindAllProductController(ctx, usecaseproduct.NewProductFindAll(productRepo))
		p.GET("/", findAllController.GetAllProducts)

		findByIDController := controllerproduct.NewFindProductByIDController(ctx, usecaseproduct.NewProductFindByID(productRepo))
		p.GET("/:id", findByIDController.GetProductByID)

		findByCategoryController := controllerproduct.NewFindProductByCategoryController(ctx, usecaseproduct.NewProductFindByCategory(productRepo))
		p.GET("/category/:id", findByCategoryController.GetProductByCategory)

		updateController := controllerproduct.NewUpdateProductController(ctx, usecaseproduct.NewProductUpdate(productRepo))
		p.PUT("/:id", updateController.PutUpdateProduct)

	}

	o := s.router.Group("/kitchencontrol/api/v1/orders")
	{

		orderController := controllerorder.NewCreateOrderController(ctx, usecaseorder.NewOrderCreate(orderRepo, customerRepo, productRepo))
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
