package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andrersp/test-containers/repository"
	"github.com/andrersp/test-containers/testhelpers"
	"github.com/andrersp/test-containers/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CustomerHandlerTestSuit struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	repository  *repository.Repository
	ctx         context.Context
}

func (suite *CustomerHandlerTestSuit) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := testhelpers.CratePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	suite.pgContainer = pgContainer
	repository, err := repository.NewRepository(suite.ctx, suite.pgContainer.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	suite.repository = repository
}

func (s *CustomerHandlerTestSuit) TearDownSuite() {
	if err := s.pgContainer.Terminate(s.ctx); err != nil {
		log.Fatal()
	}
}

func (s *CustomerHandlerTestSuit) TestCreateCustomer() {
	t := s.T()

	createUseCase := usecase.NewCreateCustomerUseCase(s.repository)

	payload := `{"name":"Jon Snow", "email": "email@mail.com"}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewCreateCustomerUseCase(createUseCase)
	if assert.NoError(t, h.Execute(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}

}

func (s *CustomerHandlerTestSuit) TestGetCustomer() {
	t := s.T()

	result, err := s.repository.GetCustomerByEmail(s.ctx, "john@gmail.com")
	assert.NoError(t, err)
	assert.Equal(t, "John", result.Name)
}

func TestCustomerHandlerTestSuit(t *testing.T) {
	suite.Run(t, new(CustomerHandlerTestSuit))
}
