test:
	go test -v -cover ./...

tidy:
	go mod tidy	
	
run:
	go run cmd/kitchencontrol/main.go	