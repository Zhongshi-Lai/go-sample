package sample

import (
	"context"
	sampleV1 "go-sample/api_gen/sample/v1"
)

func (s *Service) Echo(context.Context, *sampleV1.StringMessage) (*sampleV1.StringMessage, error) {
	return &sampleV1.StringMessage{
		Value: "hellov1",
	}, nil
}
