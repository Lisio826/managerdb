package test

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"managerdb/logger"
	"testing"
)

func TestTestfa(t *testing.T) {
	//MainLogger := logger.NewLogger("logs/test.log", zapcore.DebugLevel, 1, 2, 7, true, "Main").Sugar()
	mainLogger := logger.NewLogger("logs/test.log", zapcore.DebugLevel, 1, 3, 7, true, "Main").Sugar()
	for {
		mainLogger.Info("PanicError==》》--" + fmt.Sprint("aaaaaaaaaaaaaaaaaaaa"))
		mainLogger.Debug("i am debug",zap.String("key","debug"))
		mainLogger.Debug("aaaaaaa","12312312")
	}
}
