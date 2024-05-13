package main

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/server"
)

func main() {
	server := server.NewServer()

	server.Initialization().Run(":8080")

}
