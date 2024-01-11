package logger

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogConf struct {
	Name  string
	Path  string
	Debug bool
}

var CustomerLogger *zap.Logger

func InitGlobalLogger(name, path string, debug bool) *zap.Logger {

	if name == "" {
		panic("logger must have name")
	}

	if path == "" {
		panic("logger must have storage path")
	}

	// LoggerEncoderConfig
	fileLoggerEncoderConfig := zap.NewProductionEncoderConfig()
	fileLoggerEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleLoggerEncoderConfig := zap.NewProductionEncoderConfig()
	consoleLoggerEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleLoggerEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// createEncoder
	fileEncoder := zapcore.NewJSONEncoder(fileLoggerEncoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(consoleLoggerEncoderConfig)

	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}

	writeSyncer := zapcore.AddSync(lumberJackLogger)

	// 开启文件及行号
	caller := zap.AddCaller()

	// error以上级别才开启Stacktrace
	stackTrace := zap.AddStacktrace(zap.ErrorLevel)

	// 创建logger
	//loggerLevel := zapcore.Level(zap.DebugLevel)

	var core zapcore.Core

	// 生产环境全面使用k8s,不再使用supervisor的  stdout>>xxx.log落到文件这种形式
	// k8s的标准输出,是落到pod内存中的,除非你额外要求永久存储
	// 所以主要使用 fileEncoder 落到文件上,再从文件上传到阿里云
	// 标准输出在k8s的这个趋势下,也仅用于输出到屏幕上看,而非落到某个xxx.log文件了
	if debug {
		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writeSyncer, zap.DebugLevel),
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(fileEncoder, writeSyncer, zap.InfoLevel)
	}

	if core == nil {
		panic("log core 初始化失败")
	}

	logger := zap.New(core, caller, stackTrace)
	if logger == nil {
		panic("初始化logger失败: " + name)
	}

	zap.ReplaceGlobals(logger)

	CustomerLogger = logger

	return logger
}

func New(conf *LogConf) {
	InitGlobalLogger(conf.Name, conf.Path, conf.Debug)
	zap.L().Sugar().Info("hello world my zap logger")
	zap.L().Sugar().Error("this will be an error")

}

func WithContext(ctx context.Context) *zap.Logger {
	if CustomerLogger == nil {
		panic("please init CustomerLogger first")
	}
	if ctx == nil {
		return CustomerLogger
	}

	// get field from ctx and

	//if md, ok := metadata.FromContext(ctx); ok {
	//	if li, ok := md[loggerKey]; ok {
	//		if l, ok := li.(*zap.Logger); ok {
	//			return l
	//		}
	//	}
	//}
	return CustomerLogger
}
