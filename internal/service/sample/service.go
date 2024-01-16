package sample

import (
	sampleV1 "go-sample/api_gen/sample/v1"
	"go-sample/pkg"
)

type YourServiceServer struct {
	Tools *pkg.Tools
	sampleV1.UnimplementedYourServiceServer
}

func NewYourServiceServer(tools *pkg.Tools) *YourServiceServer {
	return &YourServiceServer{Tools: tools}
}
