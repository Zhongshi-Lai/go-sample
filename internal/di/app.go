package di

import (
	"go-sample/pkg/server"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewApp(ginServer *gin.Engine, tools *server.Tools, grpcServer *grpc.Server) (*server.App, error) {
	return &server.App{
		GinServer:  ginServer,
		Tools:      tools,
		GRPCServer: grpcServer,
	}, nil
}
