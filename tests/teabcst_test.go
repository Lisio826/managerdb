package test

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"managerdb/logger"
	"testing"
)

func TestTestfa(t *testing.T) {
	MainLogger := log.NewLogger("logs/test.log", zapcore.DebugLevel, 1, 2, 7, true, "Main").Sugar()
	MainLogger.Info("PanicError==》》--" + fmt.Sprint("aaaaaaaaaaaaaaaaaaaa"))
	MainLogger.Debug("i am debug",zap.String("key","debug"))
	MainLogger.Debug("aaaaaaa","12312312")
}
