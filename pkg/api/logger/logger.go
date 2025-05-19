package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/thanhpv3380/api/utils"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

type LoggerConfig struct {
	LogFile         string // Tên tệp log
	LogLevel        string // Mức độ log: "debug", "info", "warn", "error", "fatal"
	TimestampFormat string // Định dạng timestamp
	MaxSize         int    // MB
	MaxAge          int    // Ngày giữ log
	MaxBackups      int    // Số lượng bản sao lưu
	LocalTime       bool   // Sử dụng giờ địa phương
	Compress        bool   // Nén log
	IsConsole       bool   // Ghi log ra console
}

func DefaultLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		LogFile:         "./logs/app.log",
		LogLevel:        "info",
		TimestampFormat: "2006-01-02 15:04:05.000",
		MaxSize:         10,
		MaxAge:          7,
		MaxBackups:      100,
		LocalTime:       true,
		Compress:        true,
		IsConsole:       true,
	}
}

func NewLogger(config *LoggerConfig) {
	if config == nil {
		config = DefaultLoggerConfig()
	}

	logDir := filepath.Dir(config.LogFile)
	utils.CheckExistOrMake(logDir)

	lumberjackLogger := &lumberjack.Logger{
		Filename:   config.LogFile,
		MaxSize:    config.MaxSize,
		MaxAge:     config.MaxAge,
		MaxBackups: config.MaxBackups,
		LocalTime:  config.LocalTime,
		Compress:   config.Compress,
	}

	var level zapcore.Level
	switch config.LogLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}

	// Định dạng log
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(config.TimestampFormat),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var core zapcore.Core
	if config.IsConsole {
		consoleCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			level,
		)

		fileCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(lumberjackLogger),
			level,
		)

		core = zapcore.NewTee(consoleCore, fileCore)
	} else {
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(lumberjackLogger),
			level,
		)
	}

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Logger = logger.Sugar()

	Logger.Info("Init logger successfully")
}

func Info(message string, fields ...interface{}) {
	if len(fields) == 1 {
		switch v := fields[0].(type) {
		case map[string]interface{}:
			args := utils.FlattenMap(v)
			Logger.Infow(message, args...)
			return
		case error:
			Logger.Infow(message, "error", v.Error())
			return
		default:
			Logger.Warnw("Invalid field type provided", "field", fmt.Sprintf("%v", fields[0]))
		}
	} else if len(fields)%2 == 0 {
		Logger.Infow(message, fields...)
	} else {
		Logger.Warnw("Invalid fields provided, expected key-value pairs", "fields", fmt.Sprintf("%v", fields))
	}
}

func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

func Warn(message string, fields ...interface{}) {
	if len(fields) == 1 {
		switch v := fields[0].(type) {
		case map[string]interface{}:
			args := utils.FlattenMap(v)
			Logger.Warnw(message, args...)
			return
		case error:
			Logger.Warnw(message, "error", v.Error())
			return
		default:
			Logger.Warnw("Invalid field type provided", "field", fmt.Sprintf("%v", fields[0]))
		}
	} else if len(fields)%2 == 0 {
		Logger.Warnw(message, fields...)
	} else {
		Logger.Warnw("Invalid fields provided, expected key-value pairs", "fields", fmt.Sprintf("%v", fields))
	}
}

func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}

func Error(message string, err error, fields ...interface{}) {
	if err != nil {
		Logger.Errorw(message, "error", err)
	} else {
		Logger.Error(message)
	}
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}
