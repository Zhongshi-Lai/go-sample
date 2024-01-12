package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewGinServer() *gin.Engine {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// TODO(laizhongshi): 注册路由

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := r.Run(); err != nil {
		panic("run http router error")
	}

	return r
}
