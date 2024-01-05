


swagger:
	@./tools/swagger/swag init --parseDependency --parseInternal


lint:
	@golangci-lint run -c ./tools/golangci-lint/.golangci.yml --allow-parallel-runners

fmt:
	@goimports -w .

checkDependencies:
	@go mod graph | gmchart



VERSION := $(shell git rev-parse --short HEAD)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

GOLDFLAGS += -X main.Version=$(VERSION)
GOLDFLAGS += -X main.Buildtime=$(BUILDTIME)
GOFLAGS = -ldflags "$(GOLDFLAGS)"

run: build
	./mybinary

build:
	go build -o mybinary $(GOFLAGS) .