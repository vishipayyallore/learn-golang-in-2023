package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll(status string) ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Ashish", "Bangalore", "560001", "2000-01-01", "1"},
		{"1002", "Manish", "Hyderabad", "560001", "2000-01-01", "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}
