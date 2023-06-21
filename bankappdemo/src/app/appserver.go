package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"bankappdemo/domain"
	"bankappdemo/logger"
	"bankappdemo/services"

	"github.com/gorilla/mux"
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

	// Defining the routes for Demo Customers
	muxRouter.HandleFunc("/api/democustomers", GetAllDemoCustomersHandler).Methods(http.MethodGet)
	muxRouter.HandleFunc("/api/democustomersjson", GetAllDemoCustomersInJsonHandler).Methods(http.MethodGet)
	muxRouter.HandleFunc("/api/democustomersxml", GetAllDemoCustomersInXmlHandler).Methods(http.MethodGet)

	// Wiring up the Customer Handlers
	// customerHandlers := &CustomersHandlers{customerService: services.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customersHandlersOld := &CustomersHandlersOld{customerService: services.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Defining the routes for Customers Old
	muxRouter.HandleFunc("/api/customersold", customersHandlersOld.GetAllCustomersHandler).Methods(http.MethodGet)
	muxRouter.HandleFunc("/api/customersold/{customer_id:[0-9]+}", customersHandlersOld.GetCustomer).Methods(http.MethodGet)

	customersHandlers := &CustomersHandlers{customerService: services.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// Defining the routes for Customers
	muxRouter.HandleFunc("/api/customers", customersHandlers.GetAllCustomersHandler).Methods(http.MethodGet)
	muxRouter.HandleFunc("/api/customers/{customer_id:[0-9]+}", customersHandlers.GetCustomer).Methods(http.MethodGet)

	// Starting the server
	logger.Info("Starting the Server on " + hostServer)

	log.Fatal(http.ListenAndServe(hostServer, muxRouter))

}
