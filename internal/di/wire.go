//go:build wireinject
// +build wireinject

package di

import (
	"go-sample/internal/server"
	"go-sample/internal/service"
	"go-sample/pkg"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

//type GinApp struct {
//	GinServer *gin.Engine
//}

//func newGinApp(ginServer *gin.Engine) (*GinApp, error) {
//	return &GinApp{
//		GinServer: ginServer,
//	}, nil
//}

//func InitializeGinApp() (app *GinApp, closeFunc func(), err error) {
//	wire.Build(server.NewGinServer, service.ProviderSet, service.NewAllService, pkg.NewAllTools, newGinApp)
//	return &GinApp{}, func() {}, nil
//}

type GRPCApp struct {
	GRPCServer *grpc.Server
}

func newGRPCApp(grpcServer *grpc.Server) (*GRPCApp, error) {
	return &GRPCApp{
		GRPCServer: grpcServer,
	}, nil
}

func InitializeGRPCApp() (app *GRPCApp, closeFunc func(), err error) {
	wire.Build(server.NewGRPCServer, service.ProviderSet, service.NewAllService, pkg.NewAllTools, newGRPCApp)
	return &GRPCApp{}, func() {}, nil
}
