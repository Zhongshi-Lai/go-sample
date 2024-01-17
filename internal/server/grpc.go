package server

import (
	"fmt"
	sample_v1 "go-sample/api_gen/sample/v1"
	"go-sample/internal/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func NewGRPCServer(allService *service.AllService) *grpc.Server {

	port := 8064

	fmt.Println("ready to start grpc server")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	sample_v1.RegisterSampleServiceServer(grpcServer, allService.SampleService)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	fmt.Println("finish to start grpc server")

	// run proxy

	runProxy()

	return grpcServer
}
