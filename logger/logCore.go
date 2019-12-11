package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// https://studygolang.com/articles/17394
/**
 * 获取日志
 * filePath 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 * serviceName 服务名
 */
func NewLogger(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool, serviceName string) *zap.Logger {
	core := newCore(filePath, level, maxSize, maxBackups, maxAge, compress)

	// 开启开发模式，堆栈跟踪
	//caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", serviceName))

	//// 构造日志
	//return zap.New(core, caller, development, filed)

	// 构造日志
	return zap.New(core, development, filed)
}

/**
 * zapcore构造
 */
func newCore(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.Core {
	//日志文件路径配置2
	hook := lumberjack.Logger{
		Filename:   filePath,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   compress,   // 是否压缩
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:    "time",
		LevelKey:   "level",
		NameKey:    "logger",
		CallerKey:  "lineNum",
		MessageKey: "msg",
		//StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),               // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),zapcore.AddSync(&hook)), // 打印到控制台和文件 zapcore.AddSync(os.Stdout),zapcore.AddSync(&hook)
		atomicLevel,                                         // 日志级别
	)
}

// 另一种写法
var errorLogger *zap.SugaredLogger
var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func a()  {
	fileName := "zap.log"
	level := getLoggerLevel("debug")
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   1 << 30, //1G
		LocalTime: true,
		Compress:  true,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	errorLogger = logger.Sugar()
}
//func Debug(args ...interface{}) {
//	errorLogger.Debug(args...)
//}
//
//func Debugf(template string, args ...interface{}) {
//	errorLogger.Debugf(template, args...)
//}
//
//func Info(args ...interface{}) {
//	errorLogger.Info(args...)
//}
//
//func Infof(template string, args ...interface{}) {
//	errorLogger.Infof(template, args...)
//}
//
//func Warn(args ...interface{}) {
//	errorLogger.Warn(args...)
//}
//
//func Warnf(template string, args ...interface{}) {
//	errorLogger.Warnf(template, args...)
//}
//
//func Error(args ...interface{}) {
//	errorLogger.Error(args...)
//}
//
//func Errorf(template string, args ...interface{}) {
//	errorLogger.Errorf(template, args...)
//}
//
//func DPanic(args ...interface{}) {
//	errorLogger.DPanic(args...)
//}
//
//func DPanicf(template string, args ...interface{}) {
//	errorLogger.DPanicf(template, args...)
//}
//
//func Panic(args ...interface{}) {
//	errorLogger.Panic(args...)
//}
//
//func Panicf(template string, args ...interface{}) {
//	errorLogger.Panicf(template, args...)
//}
//
//func Fatal(args ...interface{}) {
//	errorLogger.Fatal(args...)
//}
//
//func Fatalf(template string, args ...interface{}) {
//	errorLogger.Fatalf(template, args...)
//}
