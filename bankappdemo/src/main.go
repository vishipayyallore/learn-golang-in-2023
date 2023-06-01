package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {

	hostServer := "localhost:8080"

	// Defining the routes
	http.HandleFunc("/", apiRootHandler)
	http.HandleFunc("/api", apiRootHandler)

	// Defining the routes for Customers
	http.HandleFunc("/api/customers", getAllCustomersHandler)

	// Starting the server
	fmt.Println("Starting the Server on ", hostServer)

	log.Fatal(http.ListenAndServe(hostServer, nil))

}

func apiRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Bank App Web API in Go!")
}

func getAllCustomersHandler(w http.ResponseWriter, r *http.Request) {

	customers := []Customer{
		{Name: "John", City: "New York", ZipCode: "12345"},
		{Name: "Jane", City: "New York", ZipCode: "12345"},
		{Name: "Jack", City: "New York", ZipCode: "12345"},
		{Name: "Jill", City: "New York", ZipCode: "12345"},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(customers)
}
