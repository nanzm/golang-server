package logx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

var Zap *zap.Logger
var simpleLogger *zap.SugaredLogger

func Init(path string) {
	hook := lumberjack.Logger{
		Filename:   path, // 日志文件路径
		MaxSize:    100,  // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 10,   // 日志文件最多保存多少个备份
		MaxAge:     14,   // 文件最多保存多少天
		Compress:   true, // 是否压缩
	}

	colorEncoder := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	encoder := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	Debug := zap.NewAtomicLevelAt(zap.DebugLevel)
	Warn := zap.NewAtomicLevelAt(zap.WarnLevel)

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(colorEncoder), zapcore.AddSync(os.Stdout), Debug),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoder), zapcore.AddSync(&hook), Warn),
	)

	// 构造日志
	Zap = zap.New(core, zap.AddCaller())
	simpleLogger = Zap.WithOptions(zap.AddCallerSkip(1)).Sugar()
}

func getSimpleLogger() *zap.SugaredLogger {
	if simpleLogger == nil {
		log.Fatal("please init before use logx")
		return nil
	}
	return simpleLogger
}

// 兼容
func Println(args ...interface{}) {
	getSimpleLogger().Info(args...)
}

// 兼容
func Printf(template string, args ...interface{}) {
	getSimpleLogger().Infof(template, args...)
}

// debug
func Debugf(template string, args ...interface{}) {
	getSimpleLogger().Debugf(template, args...)
}

// info
func Infof(template string, args ...interface{}) {
	getSimpleLogger().Infof(template, args...)
}

func Info(args ...interface{}) {
	getSimpleLogger().Info(args...)
}

// warn
func Warnf(template string, args ...interface{}) {
	getSimpleLogger().Warnf(template, args...)
}

func Warn(args ...interface{}) {
	getSimpleLogger().Warn(args...)
}

// error
func Error(args ...interface{}) {
	getSimpleLogger().Error(args...)
}

func Errorf(template string, args ...interface{}) {
	getSimpleLogger().Errorf(template, args...)
}

// panic
func Panic(args ...interface{}) {
	getSimpleLogger().Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	getSimpleLogger().Panicf(template, args...)
}

// fatal
func Fatal(args ...interface{}) {
	getSimpleLogger().Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	getSimpleLogger().Fatalf(template, args...)
}
