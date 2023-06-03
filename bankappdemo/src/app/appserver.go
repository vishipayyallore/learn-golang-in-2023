package app

import (
	"fmt"
	"log"
	"net/http"
)

func StartBankServer() {
	hostServer := "localhost:8080"

	serverMux := http.NewServeMux()

	// Defining the routes
	serverMux.HandleFunc("/", ApiRootHandler)
	serverMux.HandleFunc("/api", ApiRootHandler)

	// Defining the routes for Customers
	serverMux.HandleFunc("/api/customers", GetAllCustomersHandler)
	serverMux.HandleFunc("/api/customersjson", GetAllCustomersInJsonHandler)
	serverMux.HandleFunc("/api/customersxml", GetAllCustomersInXmlHandler)

	// Starting the server
	fmt.Println("Starting the Server on ", hostServer)
	log.Fatal(http.ListenAndServe(hostServer, serverMux))

}
