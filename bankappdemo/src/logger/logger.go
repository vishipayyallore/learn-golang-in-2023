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
