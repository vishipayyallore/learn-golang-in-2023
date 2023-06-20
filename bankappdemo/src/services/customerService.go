package services

import (
	"bankappdemo/domain"
	"bankappdemo/dtos"
	"bankappdemo/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errs.AppError)

	GetCustomer(string) (*dtos.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (service DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	return service.repository.FindAll(status)
}

func (service DefaultCustomerService) GetCustomer(id string) (*dtos.CustomerResponse, *errs.AppError) {
	c, err := service.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	response := dtos.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.Status,
	}

	return &response, nil
}

func NewCustomerService(customerRepository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: customerRepository}
}
