package server

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Tools struct {
}

type App struct {
	GinServer  *gin.Engine
	Tools      *Tools
	GRPCServer *grpc.Server
}
