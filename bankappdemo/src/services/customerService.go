package services

import (
	"bankappdemo/domain"
	"bankappdemo/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, *errs.AppError)

	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (service DefaultCustomerService) GetAllCustomers() ([]domain.Customer, *errs.AppError) {
	return service.repository.FindAll()
}

func (service DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return service.repository.FindById(id)
}

func NewCustomerService(customerRepository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: customerRepository}
}
