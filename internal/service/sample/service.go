package sample

import (
	sampleV1 "go-sample/api_gen/sample/v1"
	"go-sample/pkg"
)

type Service struct {
	Tools *pkg.Tools
	sampleV1.UnimplementedSampleServiceServer
}

func NewService(tools *pkg.Tools) *Service {
	return &Service{Tools: tools}
}
