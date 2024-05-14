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