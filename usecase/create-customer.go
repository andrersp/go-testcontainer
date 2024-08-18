package usecase

import (
	"context"
	"errors"

	"github.com/andrersp/test-containers/customer"
	"github.com/andrersp/test-containers/repository"
)

type CreateCustomerUsecase interface {
	Execute(customer.Customer) (customer.Customer, error)
}

type CreateCustomer struct {
	customerRepository *repository.Repository
}

// Execute implements CreateCustomerUsecase.
func (c *CreateCustomer) Execute(customer customer.Customer) (customer.Customer, error) {
	ctx := context.Background()

	existUser, err := c.customerRepository.GetCustomerByEmail(ctx, customer.Email)
	if err == nil {
		return existUser, errors.New("Duplicate")
	}
	return c.customerRepository.CreateCustomer(ctx, customer)
}

func NewCreateCustomerUseCase(repository *repository.Repository) CreateCustomerUsecase {
	return &CreateCustomer{customerRepository: repository}
}
