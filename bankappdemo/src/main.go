package main

import (
	"bankappdemo/app"
	"bankappdemo/logger"
	"fmt"
)

func main() {
	logger.InitializeLogger()

	if logger.Logger == nil {
		fmt.Println("logger.Logger is nil")
	}

	logger.Logger.Info("Starting appserver::StartBankServer()")

	app.StartBankServer()
}
