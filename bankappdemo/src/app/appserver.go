package app

import (
	"fmt"
	"log"
	"net/http"

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

	// Defining the routes for Customers
	muxRouter.HandleFunc("/api/customersjson", GetAllCustomersInJsonHandler).Methods(http.MethodGet)
	muxRouter.HandleFunc("/api/customersxml", GetAllCustomersInXmlHandler).Methods(http.MethodGet)
	muxRouter.HandleFunc("/api/customers", GetAllCustomersHandler).Methods(http.MethodGet)
	muxRouter.HandleFunc("/api/customers/{customer_id:[0-9]+}", GetCustomerByIdHandler).Methods(http.MethodGet)
	muxRouter.HandleFunc("/api/customers", CreateCustomerHandler).Methods(http.MethodPost)

	// Starting the server
	fmt.Println("Starting the Server on ", hostServer)
	log.Fatal(http.ListenAndServe(hostServer, muxRouter))

}

func GetCustomerByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Customer ID: %v\n", vars["customer_id"])
}

func CreateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Customer Handler\n")
}
