package main

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(logPath, logLevel string) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   logPath, // 日志路径
		MaxSize:    1024,    // MB
		MaxAge:     7,       // days
		MaxBackups: 3,
		Compress:   true, //是否压缩，默认为false
	}
	syncer := zapcore.AddSync(&hook)

	var level zapcore.Level
	switch strings.ToLower(logLevel) {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(config), syncer, level)

	logger := zap.New(core)
	logger.Info("logger init succeed")
	return logger
}

func main() {
	logger := InitLogger("/tmp/test.log", "info")
	logger.Info("info log", zap.Int("line", 46))
	logger.Warn("warn log", zap.Int("line", 47))
	logger.Error("error log", zap.Int("line", 48))
	fmt.Println("done: ", time.Now().Format(time.RFC3339))
}
