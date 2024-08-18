package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andrersp/test-containers/repository"
	"github.com/andrersp/test-containers/testhelpers"
	"github.com/andrersp/test-containers/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreaCustomer(t *testing.T) {

	ctx := context.Background()
	pgContainer, err := testhelpers.CratePostgresContainer(ctx)
	assert.NoError(t, err)
	customerRepository, err := repository.NewRepository(ctx, pgContainer.ConnectionString)
	assert.NoError(t, err)
	crateUseCase := usecase.NewCreateCustomerUseCase(customerRepository)

	payload := `{"name":"Jon Snow", "email": "email@mail.com"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewCreateCustomerUseCase(crateUseCase)
	if assert.NoError(t, h.Execute(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}

}

func TestCreaCustomerDuplicate(t *testing.T) {

	ctx := context.Background()
	pgContainer, err := testhelpers.CratePostgresContainer(ctx)
	assert.NoError(t, err)
	customerRepository, err := repository.NewRepository(ctx, pgContainer.ConnectionString)
	assert.NoError(t, err)
	crateUseCase := usecase.NewCreateCustomerUseCase(customerRepository)

	payload := `{"name":"Jon Snow", "email": "john@gmail.com"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewCreateCustomerUseCase(crateUseCase)
	if assert.NoError(t, h.Execute(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

}
