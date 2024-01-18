#swagger:
#	@./tools/swagger/swag init --parseDependency --parseInternal


#lint:
#	@golangci-lint run -c ./tools/golangci-lint/.golangci.yml --allow-parallel-runners
#
#fmt:
#	@goimports -w .
#
#checkDependencies:
#	@go mod graph | gmchart






#
##run: build
##	./mybinary
#


#runserver:
#	go run cmd/classic_jin_http_server/main.go --conf=./config/test --ginPort=8063


di:
	cd internal/di && wire

# make PROTO_FILE_NAME=./sample/v1/sample.proto proto_gen
PROTO_FILE_NAME = ""
proto_gen:
	cd proto && protoc -I . \
        --go_out ../api_gen --go_opt paths=source_relative \
        --go-grpc_out ../api_gen --go-grpc_opt paths=source_relative \
        $(PROTO_FILE_NAME)
	cd proto && protoc -I . --grpc-gateway_out ../api_gen \
        --grpc-gateway_opt paths=source_relative \
        $(PROTO_FILE_NAME)


# build grpc server with git commit hash
VERSION := $(shell git rev-parse --short HEAD)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

GOLDFLAGS += -X main.Version=$(VERSION)
GOLDFLAGS += -X main.BuildTime=$(BUILDTIME)
GOFLAGS = -ldflags "$(GOLDFLAGS)"

build_grpc:
	cd cmd/mixed_grpc_http_server && go build -o ../../build/cmd/mixed_grpc_http_server/ $(GOFLAGS) .
