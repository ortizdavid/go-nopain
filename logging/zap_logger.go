package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)


func NewZapLogger(logRootPath string, logFileName string, maxSize int, maxBackups int, maxAge int) *zap.Logger {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logRootPath +"/"+logFileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   true,
	}
	// Create a zap core that writes logs to the lumberjack logger
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(lumberjackLogger),
		zap.DebugLevel,
	)
	// Create a logger with the zap core
	logger := zap.New(zapCore)
	return logger
}