package logger

import (
	"go.uber.org/zap"
)

var _logger *zap.Logger

func init() {
	var err error

	_logger, err = zap.NewProduction(zap.AddCallerSkip(1))

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
