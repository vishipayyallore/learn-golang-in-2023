package services

import (
	"bankappdemo/domain"
)

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (service DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, error) {
	return service.repo.FindAll(status)
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}
