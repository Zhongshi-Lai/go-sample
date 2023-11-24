swagger:
	@./swag init
	@go run main.go

lint:
	@./golangci-lint run -c .golangci.yml --allow-parallel-runners
