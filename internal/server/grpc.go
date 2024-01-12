package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func NewGRPCServer() *grpc.Server {

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8064))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	//sample_v1.RegisterYourServiceServer(grpcServer, sample.Echo)

	grpcServer.Serve(lis)

	return grpcServer
}
