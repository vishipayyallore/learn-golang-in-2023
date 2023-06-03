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

	// Defining the routes for Customers
	muxRouter.HandleFunc("/api/customers", GetAllCustomersHandler)
	muxRouter.HandleFunc("/api/customersjson", GetAllCustomersInJsonHandler)
	muxRouter.HandleFunc("/api/customersxml", GetAllCustomersInXmlHandler)

	// Starting the server
	fmt.Println("Starting the Server on ", hostServer)
	log.Fatal(http.ListenAndServe(hostServer, muxRouter))

}
