package app

import (
	"fmt"
	"log"
	"net/http"
)

func startBankServer() {
	hostServer := "localhost:8080"

	// Defining the routes
	http.HandleFunc("/", apiRootHandler)
	http.HandleFunc("/api", apiRootHandler)

	// Defining the routes for Customers
	http.HandleFunc("/api/customers", getAllCustomersHandler)
	http.HandleFunc("/api/customersjson", getAllCustomersInJsonHandler)
	http.HandleFunc("/api/customersxml", getAllCustomersInXmlHandler)

	// Starting the server
	fmt.Println("Starting the Server on ", hostServer)
	log.Fatal(http.ListenAndServe(hostServer, nil))

}
