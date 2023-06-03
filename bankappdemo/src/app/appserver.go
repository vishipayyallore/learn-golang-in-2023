package app

import (
	"fmt"
	"log"
	"net/http"
)

func StartBankServer() {
	hostServer := "localhost:8080"

	newServerMux := http.NewServeMux()

	// Defining the routes
	newServerMux.HandleFunc("/", ApiRootHandler)
	newServerMux.HandleFunc("/api", ApiRootHandler)

	// Defining the routes for Customers
	newServerMux.HandleFunc("/api/customers", GetAllCustomersHandler)
	newServerMux.HandleFunc("/api/customersjson", GetAllCustomersInJsonHandler)
	newServerMux.HandleFunc("/api/customersxml", GetAllCustomersInXmlHandler)

	// Starting the server
	fmt.Println("Starting the Server on ", hostServer)
	log.Fatal(http.ListenAndServe(hostServer, newServerMux))

}
