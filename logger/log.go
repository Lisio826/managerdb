package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"managerdb/conf"
)

var (
	mainLogger *zap.SugaredLogger
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
	mainLogger = NewLogger(managerPath, zapcore.DebugLevel, 1, 2, 7, true, "Main").Sugar()
}

func Debug(args ...interface{}) {
	mainLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	mainLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	mainLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	mainLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	mainLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	mainLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	mainLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	mainLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	mainLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	mainLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	mainLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	mainLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	mainLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	mainLogger.Fatalf(template, args...)
}
