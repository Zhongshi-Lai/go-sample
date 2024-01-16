//go:build wireinject
// +build wireinject

package di

import (
	"go-sample/internal/server"
	"go-sample/internal/service"
	"go-sample/pkg"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

type App struct {
	GinServer  *gin.Engine
	GRPCServer *grpc.Server
}

func newApp(ginServer *gin.Engine, grpcServer *grpc.Server) (*App, error) {
	return &App{
		GinServer:  ginServer,
		GRPCServer: grpcServer,
	}, nil
}

func InitializeApp() (app *App, closeFunc func(), err error) {
	wire.Build(server.ProviderSet, service.ProviderSet, service.NewAllService, pkg.NewAllTools, newApp)
	return &App{}, func() {}, nil
}
