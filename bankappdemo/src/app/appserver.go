package app

import (
	"fmt"
	"log"
	"net/http"
)

func StartBankServer() {
	hostServer := "localhost:8080"

	// Defining the routes
	http.HandleFunc("/", ApiRootHandler)
	http.HandleFunc("/api", ApiRootHandler)

	// Defining the routes for Customers
	http.HandleFunc("/api/customers", GetAllCustomersHandler)
	http.HandleFunc("/api/customersjson", GetAllCustomersInJsonHandler)
	http.HandleFunc("/api/customersxml", GetAllCustomersInXmlHandler)

	// Starting the server
	fmt.Println("Starting the Server on ", hostServer)
	log.Fatal(http.ListenAndServe(hostServer, nil))

}

type CustomerV2 struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}
