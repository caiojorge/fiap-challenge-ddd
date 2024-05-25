package main

import (
	"github.com/caiojorge/fiap-challenge-ddd/cmd/kitchencontrol/server"
	"github.com/caiojorge/fiap-challenge-ddd/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	server := server.NewServer()

	server.Initialization()

	docs.SwaggerInfo.BasePath = "/kitchencontrol/api/v1"
	server.GetRouter().GET("/kitchencontrol/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	server.Run(":8080")

}
