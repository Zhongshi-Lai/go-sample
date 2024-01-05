package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	confPath string
	ginPort  int64
)

func main() {

	// read config path from cmd flag
	flag.StringVar(&confPath, "conf", "", "default config path")
	flag.Int64Var(&ginPort, "ginPort", 0, "assign gin http server port")

	flag.Parse()

	// read port from

	if confPath == "" {
		panic("empty conf path")
	}

	if ginPort == 0 {
		panic("empty gin port")
	}

	viper.AddConfigPath(confPath)
	viper.SetConfigName("server")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic("file not found")
		} else {
			// Config file was found but another error was produced
			panic(err)
		}
	}

	fmt.Println(viper.AllKeys())
	fmt.Println(viper.Get("http.port"))

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
