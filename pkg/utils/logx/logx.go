package logx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Zap *zap.Logger
var simpleLogger *zap.SugaredLogger

func Init(path string) {
	hook := lumberjack.Logger{
		Filename:   path, // 日志文件路径
		MaxSize:    100,          // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 10,           // 日志文件最多保存多少个备份
		MaxAge:     14,           // 文件最多保存多少天
		Compress:   true,         // 是否压缩
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

// 兼容
func Println(args ...interface{}) {
	simpleLogger.Info(args...)
}

// 兼容
func Printf(template string, args ...interface{}) {
	simpleLogger.Infof(template, args...)
}

// debug
func Debugf(template string, args ...interface{}) {
	simpleLogger.Debugf(template, args...)
}

// info
func Infof(template string, args ...interface{}) {
	simpleLogger.Infof(template, args...)
}

func Info(args ...interface{}) {
	simpleLogger.Info(args...)
}

// warn
func Warnf(template string, args ...interface{}) {
	simpleLogger.Warnf(template, args...)
}

func Warn(args ...interface{}) {
	simpleLogger.Warn(args...)
}

// error
func Error(args ...interface{}) {
	simpleLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	simpleLogger.Errorf(template, args...)
}

// panic
func Panic(args ...interface{}) {
	simpleLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	simpleLogger.Panicf(template, args...)
}

// fatal
func Fatal(args ...interface{}) {
	simpleLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	simpleLogger.Fatalf(template, args...)
}
