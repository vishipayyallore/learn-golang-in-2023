package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _logger *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	_logger, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}

	defer _logger.Sync()

	_logger.Info("Logger initialized")
}

func Info(message string, fields ...zap.Field) {
	_logger.Info(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	_logger.Fatal(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	_logger.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	_logger.Error(message, fields...)
}
