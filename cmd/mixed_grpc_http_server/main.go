package main

import (
	"go-sample/internal/di"
	"go-sample/pkg/conf"
	"go-sample/pkg/logger"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/spf13/viper"
)

func main() {
	// init conf from
	confFilePath := os.Getenv("GO-SAMPLE-CONF")
	if confFilePath == "" {
		panic("empty config path from os env")
	}

	conf.AppConfInit(confFilePath)

	// init logger
	logger.New(&logger.LogConf{
		Name:  viper.GetString("serverLog.name"),
		Path:  viper.GetString("serverLog.path"),
		Debug: viper.GetBool("serverMode.debug"),
	})

	// di init
	// in di init ,you should start the server

	_, closeFunc, err := di.InitializeGRPCApp()
	if err != nil {
		panic(err)
	}

	closeFunc()

	// all close func
	// like mysql-pool

	// set max cpu
	runtime.GOMAXPROCS(runtime.NumCPU())

	sysSignalChan := make(chan os.Signal, 1)
	signal.Notify(sysSignalChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	for {
		s := <-sysSignalChan
		//zlog.Log.Sugar().Infof("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//closeFunc()
			//zlog.Log.Sugar().Infof("%s exit", AppName)
			//time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

}
