package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type CustomerV1 struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

func GetAllDemoCustomersHandler(w http.ResponseWriter, r *http.Request) {

	customers := GetDummyCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func GetAllDemoCustomersInJsonHandler(w http.ResponseWriter, r *http.Request) {

	customers := GetDummyCustomers()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(customers)
}

func GetAllDemoCustomersInXmlHandler(w http.ResponseWriter, r *http.Request) {

	customers := GetDummyCustomers()

	w.Header().Set("Content-Type", "application/xml")

	xml.NewEncoder(w).Encode(customers)
}

func GetDummyCustomers() []CustomerV1 {
	return []CustomerV1{
		{Name: "John", City: "New York", Zipcode: "12345"},
		{Name: "Jane", City: "New York", Zipcode: "12345"},
		{Name: "Jack", City: "New York", Zipcode: "12345"},
		{Name: "Jill", City: "New York", Zipcode: "12345"},
	}
}
