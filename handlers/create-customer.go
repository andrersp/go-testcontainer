package handlers

import (
	"net/http"

	"github.com/andrersp/test-containers/customer"
	"github.com/andrersp/test-containers/usecase"
	"github.com/labstack/echo/v4"
)

type createCustomerHandlers struct {
	useCase usecase.CreateCustomerUsecase
}

func (c *createCustomerHandlers) Execute(e echo.Context) error {
	var payload customer.Customer
	if err := e.Bind(&payload); err != nil {
		return err
	}
	response, err := c.useCase.Execute(payload)
	if err != nil {

		return e.JSON(400, err)
	}

	return e.JSON(http.StatusCreated, response)

}

func NewCreateCustomerUseCase(usecase usecase.CreateCustomerUsecase) *createCustomerHandlers {
	return &createCustomerHandlers{useCase: usecase}
}
