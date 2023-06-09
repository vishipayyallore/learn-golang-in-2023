package main

import (
	"bankappdemo/app"
	"bankappdemo/logger"
	"fmt"
)

func main() {
	if logger.Logger == nil {
		fmt.Println("logger.Logger is nil")
	}

	logger.log.Info("Starting appserver::StartBankServer()")

	app.StartBankServer()
}
