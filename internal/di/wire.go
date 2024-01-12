//go:build wireinject
// +build wireinject

package di

import (
	"go-sample/internal/server"
	pkgServer "go-sample/pkg/server"

	"github.com/google/wire"
)

func InitializeApp() (app *pkgServer.App, closeFunc func(), err error) {
	wire.Build(server.ProviderSet, NewTools, NewApp)
	return &pkgServer.App{}, func() {}, nil
}
