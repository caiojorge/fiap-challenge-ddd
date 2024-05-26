package main

import (
	"log"

	"github.com/caiojorge/fiap-challenge-ddd/docs"
	infra "github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/db"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/server"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @title Fiap Challenge DDD API
// @version 1.0
// @description This is fiap ddd challenge project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /kitchencontrol/api/v1

func main() {

	db := setupDB()
	server := server.NewServer(db)

	server.Initialization()

	// Migrate the schema
	if err := server.GetDB().AutoMigrate(&model.Customer{}, &model.Product{}, &model.Order{}, &model.OrderItem{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	docs.SwaggerInfo.BasePath = "/kitchencontrol/api/v1"
	server.GetRouter().GET("/kitchencontrol/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	server.Run(":8080")

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

	return connection
}
