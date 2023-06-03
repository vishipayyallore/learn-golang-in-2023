package services

import (
	"bankappdemo/domain"
)

type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (service DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, error) {
	return service.repository.FindAll(status)
}

func NewCustomerService(customerRepository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: customerRepository}
}
