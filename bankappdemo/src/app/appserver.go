package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"bankappdemo/domain"
	"bankappdemo/logger"
	"bankappdemo/services"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		logger.Error("Environment variable not defined...")
		os.Exit(1)
	}
}

func StartBankServer() {

	sanityCheck()

	serverAddress := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	hostServer := fmt.Sprintf("%s:%s", serverAddress, serverPort)

	// muxRouter := http.NewServeMux()
	muxRouter := mux.NewRouter()

	// Defining the routes
	muxRouter.HandleFunc("/", ApiRootHandler).Methods(http.MethodGet)
	muxRouter.HandleFunc("/api", ApiRootHandler).Methods(http.MethodGet)

	// Wiring up the Customer Handlers
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	customerService := services.NewCustomerService(customerRepositoryDb)
	customersHandlers := &CustomersHandlers{customerService}

	// Defining the routes for Customers
	muxRouter.HandleFunc("/api/customers", customersHandlers.GetAllCustomersHandler)
				.Methods(http.MethodGet)
				.Name("GetAllCustomers")

	muxRouter.HandleFunc("/api/customers/{customer_id:[0-9]+}", customersHandlers.GetCustomer)
				.Methods(http.MethodGet)
				.Name("GetCustomer")

	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	accountService := services.NewAccountService(accountRepositoryDb)
	accountHandler := AccountHandler{accountService}

	muxRouter.HandleFunc("/api/customers/{customer_id:[0-9]+}/account", accountHandler.NewAccount)
				.Methods(http.MethodPost)
				.Name("NewAccount")

	muxRouter.HandleFunc("/api/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", accountHandler.MakeTransaction)
				.Methods(http.MethodPost)
				.Name("MakeTransaction")
				
	// Starting the server
	logger.Info("Starting the Server on " + hostServer)

	log.Fatal(http.ListenAndServe(hostServer, muxRouter))

}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
