package services

import (
	"bankappdemo/domain"
)

type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (service DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, error) {
	return service.repo.FindAll(status)
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}
