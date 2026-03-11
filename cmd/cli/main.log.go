package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infow("Info log", "line", 1)
	// sugar.Errorw("Error log", "line", 2)

	// logger := zap.NewExample()
	// logger.Info("Info log", zap.Int("line", 1))
	// logger.Error("Error log", zap.Int("line", 2))
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
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

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./logs/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
