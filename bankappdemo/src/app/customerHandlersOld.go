package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"bankappdemo/services"

	"github.com/gorilla/mux"
)

type CustomersHandlersOld struct {
	customerService services.CustomerService
}

func (ch *CustomersHandlersOld) GetAllCustomersHandler(w http.ResponseWriter, r *http.Request) {

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

func (ch *CustomersHandlersOld) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	customer, err := ch.customerService.GetCustomer(customerId)

	if err != nil {
		writeResponseOld(w, r, err.Code, err.AsMessage())
		return
	} else {
		writeResponseOld(w, r, http.StatusOK, customer)
	}
}

func writeResponseOld(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
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
