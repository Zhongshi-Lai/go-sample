
swagger:
	@./tools/swagger/swag init --parseDependency --parseInternal


lint:
	@golangci-lint run -c ./tools/golangci-lint/.golangci.yml --allow-parallel-runners

fmt:
	@goimports -w .

checkDependencies:
	@go mod graph | gmchart
