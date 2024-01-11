package di

import "github.com/gin-gonic/gin"

type GinServerApp struct {
	GinServer *gin.Engine
	Tools     *Tools
}

func NewGinServerApp(ginServer *gin.Engine, tools *Tools) (*GinServerApp, error) {
	return &GinServerApp{
		GinServer: ginServer,
		Tools:     tools,
	}, nil
}
