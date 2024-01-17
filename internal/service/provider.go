package service

import (
	"go-sample/internal/service/sample"

	"github.com/google/wire"
)

// ProviderSet 初始化每个单独的sample.Service
var ProviderSet = wire.NewSet(sample.NewService)

// AllService 提供给grpcserver和httpserver进行路由注册
type AllService struct {
	SampleService *sample.Service
}

func NewAllService(sampleService *sample.Service) *AllService {
	return &AllService{SampleService: sampleService}
}
