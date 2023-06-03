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
	muxRouter.HandleFunc("/", ApiRootHandler)
	muxRouter.HandleFunc("/api", ApiRootHandler)

	// Defining the routes for Demo Customers
	muxRouter.HandleFunc("/api/democustomers", GetAllDemoCustomersHandler)
	muxRouter.HandleFunc("/api/democustomersjson", GetAllDemoCustomersInJsonHandler)
	muxRouter.HandleFunc("/api/democustomersxml", GetAllDemoCustomersInXmlHandler)

	// Defining the routes for Customers
	muxRouter.HandleFunc("/api/customersjson", GetAllCustomersInJsonHandler)
	muxRouter.HandleFunc("/api/customersxml", GetAllCustomersInXmlHandler)
	muxRouter.HandleFunc("/api/customers", GetAllCustomersHandler)
	muxRouter.HandleFunc("/api/customers/{customer_id}", GetCustomerByIdHandler)

	// Starting the server
	fmt.Println("Starting the Server on ", hostServer)
	log.Fatal(http.ListenAndServe(hostServer, muxRouter))

}

func GetCustomerByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Customer ID: %v\n", vars["customer_id"])
}
