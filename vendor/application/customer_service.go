package application

import (
	"errors"
	"time"

	"github.com/myvendor/myapp/domain"
)

type CustomerService struct {
	customerRepo domain.CustomerRepository
}

func NewCustomerService(repo domain.CustomerRepository) *CustomerService {
	return &CustomerService{repo}
}

func (cs *CustomerService) GetCustomerByID(id int) (*domain.Customer, error) {
	return cs.customerRepo.GetByID(id)
}

func (cs *CustomerService) CreateCustomer(name string, email string) (*domain.Customer, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email cannot be empty")
	}

	customer := &domain.Customer{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := cs.customerRepo.Save(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (cs *CustomerService) UpdateCustomer(customer *domain.Customer) error {
	if customer == nil {
		return errors.New("customer cannot be nil")
	}

	customer.UpdatedAt = time.Now()

	return cs.customerRepo.Update(customer)
}

func (cs *CustomerService) DeleteCustomerByID(id int) error {
	return cs.customerRepo.DeleteByID(id)
}
