package main

import (
	"bankappdemo/app"
	"bankappdemo/logger"
)

func main() {
	logger.Info("Starting appserver::StartBankServer()")

	app.StartBankServer()
}
