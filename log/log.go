package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var MainLogger *zap.Logger
var GatewayLogger *zap.Logger
var RedisLogger *zap.Logger
var LmdbLogger *zap.Logger
var HttpLogger *zap.Logger

func init() {
	MainLogger = NewLogger("./logs/manager.log", zapcore.DebugLevel, 1, 2, 7, true, "Main")
	//GatewayLogger = NewLogger("./logs/gateway.log", zapcore.DebugLevel, 128, 30, 7, true, "Gateway")
}

func Debug(msg string) {
	MainLogger.Debug(msg)
}

func Info(msg string) {
	MainLogger.Info(msg)
}

func Error(msg string) {
	MainLogger.Error(msg)
}

func Warn(msg string) {
	MainLogger.Warn(msg)
}

func Fatal(msg string) {
	MainLogger.Fatal(msg)
}
