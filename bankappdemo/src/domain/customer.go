package domain

import (
	"bankappdemo/dtos"
	"bankappdemo/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {

	// Status is a string, but it is not a good idea to use string as a type for status.
	// 1 == active, 0 == inactive, "" == All Customers
	FindAll(string) ([]Customer, *errs.AppError)

	FindById(string) (*Customer, *errs.AppError)
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dtos.CustomerResponse {
	return dtos.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}
