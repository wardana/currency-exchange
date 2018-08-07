package controller

import (
	"github.com/labstack/echo"
	"github.com/wardana/currency-exchange/helper"
	"github.com/wardana/currency-exchange/models"
	"github.com/wardana/currency-exchange/services"
)

//Controller initialize controller class
type Controller struct {
	CurrencyServices services.CurrencyInterface
	Helper           helper.Interface
}

//Interface is an interface for controller class
type Interface interface {
	Create(c echo.Context) error
	FindAll(c echo.Context) error
	UpdateCurrency(c echo.Context) error
	DeleteCurrency(c echo.Context) error
}

//Create is a handler for create new currency
func (ctr *Controller) Create(c echo.Context) error {

	payload := models.Currency{}

	// Throw err response when payload is not valid
	if err := c.Bind(&payload); err != nil {
		return ctr.Helper.HTTPBadRequest(c, "invalid payload")
	}

	data, err := ctr.CurrencyServices.Create(payload)
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPCreated(c, data)
}

// FindAll is a handler for get all available currency
func (ctr *Controller) FindAll(c echo.Context) error {

	data, err := ctr.CurrencyServices.FindAll()
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPSuccess(c, data)
}

// UpdateCurrency is a handler for get all available currency
func (ctr *Controller) UpdateCurrency(c echo.Context) error {

	payload := models.Currency{}

	id := c.Param("id")

	if id == "" {
		return ctr.Helper.HTTPBadRequest(c, "id not found")
	}

	// Throw err response when payload is not valid
	if err := c.Bind(&payload); err != nil {
		return ctr.Helper.HTTPBadRequest(c, "invalid payload")
	}

	data, err := ctr.CurrencyServices.Update(id, payload)
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPSuccess(c, data)
}

// DeleteCurrency is a handler for delete some currency data using its id
func (ctr *Controller) DeleteCurrency(c echo.Context) error {

	id := c.Param("id")

	if id == "" {
		return ctr.Helper.HTTPBadRequest(c, "id not found")
	}

	err := ctr.CurrencyServices.Delete(id)
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPSuccess(c, map[string]interface{}{})
}
