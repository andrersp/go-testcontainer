package repository

import (
	"context"
	"log"
	"testing"

	"github.com/andrersp/test-containers/customer"
	"github.com/andrersp/test-containers/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CustomerRepoTestSuit struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	repository  *Repository
	ctx         context.Context
}

func (suite *CustomerRepoTestSuit) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := testhelpers.CratePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	suite.pgContainer = pgContainer
	repository, err := NewRepository(suite.ctx, suite.pgContainer.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	suite.repository = repository
}

func (s *CustomerRepoTestSuit) TearDownSuite() {
	if err := s.pgContainer.Terminate(s.ctx); err != nil {
		log.Fatal()
	}
}

func (s *CustomerRepoTestSuit) TestCreateCustomer() {

	t := s.T()
	customer := customer.Customer{
		Name:  "Henry",
		Email: "henry@mail.com",
	}
	c, err := s.repository.CreateCustomer(s.ctx, customer)
	assert.NoError(t, err)
	assert.NotNil(t, c.Id)

}

func (s *CustomerRepoTestSuit) TestGetCustomer() {
	t := s.T()

	result, err := s.repository.GetCustomerByEmail(s.ctx, "john@gmail.com")
	assert.NoError(t, err)
	assert.Equal(t, "John", result.Name)
}

func TestCustomerRepoTestSuit(t *testing.T) {
	suite.Run(t, new(CustomerRepoTestSuit))
}
