package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func init() {
	var err error

	Logger, err = zap.NewProduction()

	if err != nil {
		panic(err)
	}

	defer Logger.Sync()

	Logger.Info("Logger initialized")
}

func Info(message string, fields ...zap.Field) {
	Logger.Info(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	Logger.Fatal(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	Logger.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	Logger.Error(message, fields...)
}
