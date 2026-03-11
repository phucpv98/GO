package logger

import (
	"go-ecommerce/pkg/setting"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	// Level Log : Debug -> Info -> Warn -> Error -> DPanic -> Panic -> Fatal
	logLevel := config.Log_level
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "dpanic":
		level = zapcore.DPanicLevel
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}
	// Encoder : JSONEncoder, ConsoleEncoder
	encoder := getEncoderLog()

	// Lumberjack Logger for log rotation
	hook := lumberjack.Logger{
		Filename:   config.File_log_name,
		MaxSize:    config.Max_size, // megabytes
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age,  //days
		Compress:   config.Compress, // disabled by default
	}

	// zapcore.NewCore(encoder, writeSyncer, levelEnabler)
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level, // log level
	)
	// logger := zap.New(core, zap.AddCaller())
	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// timestamp -> format: 2006-01-02T15:04:05.000Z0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// log level -> format: INFO, ERROR, etc.
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// caller -> format: short file path and line number
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// timestamp -> time
	encodeConfig.TimeKey = "time"

	return zapcore.NewJSONEncoder(encodeConfig)
}
