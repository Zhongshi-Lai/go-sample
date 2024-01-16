package service

import (
	"go-sample/internal/service/sample"

	"github.com/google/wire"
)

// ProviderSet 初始化每个单独的sample.Service
var ProviderSet = wire.NewSet(sample.NewYourServiceServer)

// AllService 提供给grpcserver和httpserver进行路由注册
type AllService struct {
	YourServiceServer *sample.YourServiceServer
}

func NewAllService(yourServiceServer *sample.YourServiceServer) *AllService {
	return &AllService{YourServiceServer: yourServiceServer}
}
