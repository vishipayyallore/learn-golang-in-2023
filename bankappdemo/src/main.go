package main

import (
	"bankappdemo/app"
	"bankappdemo/logger"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go!")
	logger.Log.Info("Starting the Server on ")

	app.StartBankServer()
}
