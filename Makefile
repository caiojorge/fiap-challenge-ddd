test:
	go test -v -cover ./...

tidy:
	go mod tidy	
	
run:
	go run cmd/kitchencontrol/main.go	

mysql:
	docker-compose up -d

stop:
	docker-compose down

swaggo:
	go get -u github.com/swaggo/swag/cmd/swag
	go get -u github.com/swaggo/gin-swagger
	go get -u github.com/swaggo/files
	go get -u github.com/swaggo/gin-swagger/swaggerFiles
	swag init -g cmd/kitchencontrol/main.go