
# buildinfo for all go build
# just add $(GOFLAGS) to your go build command,it will automaticly add -ldflags "-X main.BuildGitCommitHash=c5efaa7 -X main.BuildTime=2024-01-23T06:56:24Z
GITHASH := $(shell git rev-parse --short HEAD)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
GITUSER := $(shell git config user.name)
GITEMAIL := $(shell git config user.email)

GOLDFLAGS += -X main.BuildGitCommitHash=$(GITHASH)
GOLDFLAGS += -X main.BuildTime=$(BUILDTIME)
GOLDFLAGS += -X main.BuildGitUser=$(GITUSER)
GOLDFLAGS += -X main.BuildGitEmail=$(GITEMAIL)
GOFLAGS = -ldflags "$(GOLDFLAGS)"


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


build_grpc:
	cd cmd/mixed_grpc_http_server && go build -o ../../build/cmd/mixed_grpc_http_server/ $(GOFLAGS) .
