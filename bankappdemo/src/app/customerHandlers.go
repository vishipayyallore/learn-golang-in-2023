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
		if err := xml.NewEncoder(w).Encode(customers); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(customers); err != nil {
			panic(err)
		}
	}
}

func (ch *CustomersHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	customer, err := ch.customerService.GetCustomer(customerId)

	if err != nil {
		writeResponse(w, r, err.Code, err.AsMessage())
		return
	} else {
		writeResponse(w, r, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(statusCode)
		xml.NewEncoder(w).Encode(data)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(data)
	}
}
