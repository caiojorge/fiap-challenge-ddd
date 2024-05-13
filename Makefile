test:
	go test -v -cover ./...

tidy:
	go mod tidy	
	
start:
	go run cmd/kitchencontrol/main.go	