package server

import (
	"context"
	"net/http"

	sampleV1 "go-sample/api_gen/sample/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
// command-line options:
// gRPC server endpoint
// grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func runProxy() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := sampleV1.RegisterYourServiceHandlerFromEndpoint(ctx, mux, ":8064", opts)
	if err != nil {
		panic(err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)

	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}

	return
}
