//go:build wireinject
// +build wireinject

package di

import "github.com/google/wire"

func InitializeApp() (app *GinServerApp, closeFunc func(), err error) {
	wire.Build(NewAllTools, NewGinServer, NewGinServerApp)
	return &GinServerApp{}, func() {}, nil
}
