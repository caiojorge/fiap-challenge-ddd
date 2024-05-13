package main

import (
	"log"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/server"
)

func main() {
	server := server.NewServer()

	server.Initialization()

	router := server.GetRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}

}
