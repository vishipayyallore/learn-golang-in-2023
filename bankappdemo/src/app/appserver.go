package app

import (
	"fmt"
	"log"
	"net/http"

	"bankappdemo/domain"
	"bankappdemo/services"

	"github.com/gorilla/mux"
)

func StartBankServer() {
	hostServer := "localhost:8080"

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
	customerHandlers := &CustomersHandlers{customerService: services.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Defining the routes for Customers
	muxRouter.HandleFunc("/api/customers", customerHandlers.GetAllCustomersHandler).Methods(http.MethodGet)

	// Starting the server
	fmt.Println("Starting the Server on ", hostServer)
	log.Fatal(http.ListenAndServe(hostServer, muxRouter))

}
