package main

import (
	_ "github.com/caiojorge/fiap-challenge-ddd/docs"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/server"
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
// @BasePath kitchencontrol/api/v1

func main() {
	server := server.NewServer()
	//server.GetRouter().GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Initialization().Run(":8080")

}
