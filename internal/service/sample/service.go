package sample

import "go-sample/pkg/server"

type Service struct {
	App *server.App
}

func NewService(app *server.App) *Service {
	return &Service{App: app}
}
