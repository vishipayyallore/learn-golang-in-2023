package app

import (
	"encoding/json"
	"net/http"

	"bankappdemo/services"

	"github.com/gorilla/mux"
)

type CustomersHandlers struct {
	customerService services.CustomerService
}

func (ch *CustomersHandlers) GetAllCustomersHandler(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	customers, err := ch.customerService.GetAllCustomers(status)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomersHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	customer, err := ch.customerService.GetCustomer(customerId)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
