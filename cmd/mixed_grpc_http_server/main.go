package main

import (
	"flag"
	"go-sample/internal/di"
	"go-sample/pkg/logger"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/spf13/viper"
)

var (
	confPath string
	ginPort  int64
)

func readAllConf() {
	// server conf
	viper.AddConfigPath(confPath)
	viper.SetConfigName("server")
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic("conf file not found")
		} else {
			// Config file was found but another error was produced
			panic(err)
		}
	}

	// second file must use viper.MergeInConfig(); otherwise it will lose first conf content;
	// log conf
	viper.SetConfigName("log")
	viper.SetConfigType("toml")
	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic("conf file not found")
		} else {
			// Config file was found but another error was produced
			panic(err)
		}
	}
}

func main() {

	// read config path from cmd flag
	flag.StringVar(&confPath, "conf", "", "default config path")
	flag.Int64Var(&ginPort, "ginPort", 0, "assign gin http server port")

	flag.Parse()

	if confPath == "" {
		panic("empty conf path")
	}
	if ginPort == 0 {
		panic("empty gin port")
	}

	readAllConf()

	// init logger
	logger.New(&logger.LogConf{
		Name:  viper.GetString("serverLog.name"),
		Path:  viper.GetString("serverLog.path"),
		Debug: viper.GetBool("serverMode.debug"),
	})

	// di init
	// in di init ,you should start the server

	_, closeFunc, err := di.InitializeApp()
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
