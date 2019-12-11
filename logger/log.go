package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"managerdb/conf"
)

var (
	MainLogger *zap.SugaredLogger
	GatewayLogger *zap.SugaredLogger
	RedisLogger *zap.SugaredLogger
	LmdbLogger *zap.Logger
	HttpLogger *zap.Logger
	PprofLogger *zap.Logger
	MonitorLogger *zap.Logger
)

func init() {
	//PprofLogger = NewLogger("./logs/pprofLogger.log", zapcore.InfoLevel, 128, 12, 7, true, "pprof")
	//MonitorLogger = NewLogger("./logs/monitorLogger.log", zapcore.InfoLevel, 128, 12, 7, true, "monitor")
	//GatewayLogger = NewLogger("./logs/gateway.log", zapcore.DebugLevel, 128, 30, 7, true, "Gateway")
}
func InitLog()  {
	managerPath := conf.Global.LogPath.LogLocal
	MainLogger = NewLogger(managerPath, zapcore.DebugLevel, 1, 2, 7, true, "Main").Sugar()
}

func Debug(args ...interface{}) {
	MainLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	MainLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	MainLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	MainLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	MainLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	MainLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	MainLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	MainLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	MainLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	MainLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	MainLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	MainLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	MainLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	MainLogger.Fatalf(template, args...)
}
