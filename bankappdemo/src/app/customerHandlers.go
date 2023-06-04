package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"bankappdemo/services"

	"github.com/gorilla/mux"
)

type CustomersHandlers struct {
	customerService services.CustomerService
}

func (ch *CustomersHandlers) GetAllCustomersHandler(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.customerService.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomersHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	customer, err := ch.customerService.GetCustomer(customerId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
